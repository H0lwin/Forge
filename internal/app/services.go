package app

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"forge/internal/config"
	"forge/internal/domain"
	"forge/internal/generator"
	"forge/internal/generator/backend"
	"forge/internal/generator/frontend"
	"forge/internal/runner"
	"forge/internal/system"
	"forge/internal/templates"

	"github.com/charmbracelet/log"
)

type Services struct {
	Logger    *log.Logger
	FS        system.FileSystem
	Executor  system.Executor
	Config    config.Config
	Templates *templates.Engine
	Registry  *generator.Registry
}

func NewServices(stdout io.Writer, verbose bool, configPath string) (*Services, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, err
	}
	engine, err := templates.NewEngine()
	if err != nil {
		return nil, err
	}
	logger := log.NewWithOptions(stdout, log.Options{ReportTimestamp: false})
	logger.SetLevel(log.InfoLevel)
	if verbose {
		logger.SetLevel(log.DebugLevel)
	}
	fs := system.OSFileSystem{}
	ex := system.OSExecutor{}
	reg := generator.NewRegistry(
		backend.Django(fs, ex),
		backend.FastAPI(fs, ex),
		backend.Flask(fs, ex),
		backend.Express(fs, ex),
		backend.NestJS(fs, ex),
		frontend.Next(fs, ex),
		frontend.Vite(fs, ex),
	)
	return &Services{Logger: logger, FS: fs, Executor: ex, Config: cfg, Templates: engine, Registry: reg}, nil
}

func (s *Services) ValidateRequest(req domain.GenerateRequest) (domain.GenerateRequest, error) {
	if err := domain.ValidateName(req.Name); err != nil {
		return req, err
	}
	path, err := domain.ResolveProjectPath(req.BasePath, req.Name)
	if err != nil {
		return req, err
	}
	req.ProjectPath = path
	if err := domain.ValidateFrameworkExtras(req.Framework, req.Extras); err != nil {
		return req, err
	}
	return req, nil
}

func (s *Services) Generate(ctx context.Context, req domain.GenerateRequest, out io.Writer, resolver runner.FailureResolver, observer func(runner.Event)) (domain.GenerateResult, error) {
	op := "app.Generate"
	req, err := s.ValidateRequest(req)
	if err != nil {
		return domain.GenerateResult{}, domain.Wrap(op, err)
	}
	g, err := s.Registry.Get(req.Framework)
	if err != nil {
		return domain.GenerateResult{}, domain.Wrap(op, err)
	}
	if err := g.PreCheck(ctx, req); err != nil {
		return domain.GenerateResult{}, domain.Wrap(op, err, "framework", req.Framework)
	}
	baseSteps := g.Steps(ctx, req)
	extraSteps, err := s.extraSteps(req, out)
	if err != nil {
		return domain.GenerateResult{}, domain.Wrap(op, err)
	}
	steps := append(baseSteps, extraSteps...)
	start := time.Now()
	r := runner.Runner{DryRun: req.DryRun, ResolveError: resolver, Observer: observer}
	if err := r.Run(ctx, steps); err != nil {
		return domain.GenerateResult{}, domain.Wrap(op, err)
	}
	res := domain.GenerateResult{ProjectPath: req.ProjectPath, StartedAt: start, EndedAt: time.Now(), StepsTotal: len(steps), StepsDone: len(steps)}
	fmt.Fprintln(out, g.PostMessage(req, res))
	fmt.Fprintf(out, "\nBuilt in %ds (%d steps)\n", int(res.EndedAt.Sub(res.StartedAt).Seconds()), res.StepsDone)
	fmt.Fprintln(out, "Next steps:")
	fmt.Fprintf(out, "  cd %s\n", req.ProjectPath)
	switch req.Framework {
	case domain.FrameworkNext, domain.FrameworkVite, domain.FrameworkExpress, domain.FrameworkNestJS:
		fmt.Fprintln(out, "  npm install")
		fmt.Fprintln(out, "  npm run dev")
	default:
		fmt.Fprintln(out, "  python -m venv .venv")
		fmt.Fprintln(out, "  # activate venv then install dependencies")
	}
	return res, nil
}

func (s *Services) extraSteps(req domain.GenerateRequest, out io.Writer) ([]runner.Step, error) {
	steps := []runner.Step{}
	cmdOut := io.Discard
	if req.Verbose {
		cmdOut = out
	}
	writeTemplate := func(id, name, file string, data any) error {
		content, err := s.Templates.Render(name, data)
		if err != nil {
			return err
		}
		path := filepath.Join(req.ProjectPath, file)
		steps = append(steps, runner.Step{ID: id, Title: "Generating " + file, Timeout: 5 * time.Second, Action: func(context.Context) error {
			return s.FS.WriteFile(path, content, 0o644)
		}})
		return nil
	}
	if err := writeTemplate("readme", "README.md.tmpl", "README.md", map[string]string{"Name": req.Name, "RunCommand": "forge doctor"}); err != nil {
		return nil, err
	}
	if err := writeTemplate("editorconfig", "editorconfig.tmpl", ".editorconfig", map[string]string{}); err != nil {
		return nil, err
	}
	if err := writeTemplate("env", "env.tmpl", ".env", map[string]string{"Name": req.Name}); err != nil {
		return nil, err
	}
	if err := writeTemplate("env-example", "env.example.tmpl", ".env.example", map[string]string{"Name": req.Name}); err != nil {
		return nil, err
	}
	if contains(req.Extras, "ci") {
		steps = append(steps, runner.Step{ID: "ci", Title: "Creating GitHub Actions workflow", Timeout: 5 * time.Second, Action: func(context.Context) error {
			workflow := "name: ci\non: [push, pull_request]\njobs:\n  test:\n    runs-on: ubuntu-latest\n    steps:\n      - uses: actions/checkout@v4\n"
			return s.FS.WriteFile(filepath.Join(req.ProjectPath, ".github/workflows/ci.yml"), []byte(workflow), 0o644)
		}})
	}
	if contains(req.Extras, "drf") {
		steps = append(steps, runner.Step{ID: "drf", Title: "Adding Django REST Framework", Timeout: 5 * time.Second, Action: func(context.Context) error {
			p := filepath.Join(req.ProjectPath, "requirements.txt")
			b, err := appendLineIfMissing(p, "djangorestframework")
			if err != nil {
				return err
			}
			return s.FS.WriteFile(p, b, 0o644)
		}})
	}
	if contains(req.Extras, "postgres") {
		steps = append(steps, runner.Step{ID: "postgres", Title: "Configuring PostgreSQL defaults", Timeout: 5 * time.Second, Action: func(context.Context) error {
			envPath := filepath.Join(req.ProjectPath, ".env.example")
			b, err := appendMultiIfMissing(envPath, []string{
				"DB_ENGINE=postgres",
				"DB_HOST=localhost",
				"DB_PORT=5432",
				"DB_NAME=app",
				"DB_USER=app",
				"DB_PASSWORD=app",
			})
			if err != nil {
				return err
			}
			return s.FS.WriteFile(envPath, b, 0o644)
		}})
	}
	if contains(req.Extras, "pytest") {
		steps = append(steps, runner.Step{ID: "pytest", Title: "Adding pytest setup", Timeout: 5 * time.Second, Action: func(context.Context) error {
			if err := s.FS.WriteFile(filepath.Join(req.ProjectPath, "pytest.ini"), []byte("[pytest]\naddopts = -q\n"), 0o644); err != nil {
				return err
			}
			return s.FS.WriteFile(filepath.Join(req.ProjectPath, "tests", "__init__.py"), []byte(""), 0o644)
		}})
	}
	if contains(req.Extras, "precommit") {
		steps = append(steps, runner.Step{ID: "precommit", Title: "Adding pre-commit config", Timeout: 5 * time.Second, Action: func(context.Context) error {
			cfg := "repos:\n  - repo: https://github.com/pre-commit/pre-commit-hooks\n    rev: v4.6.0\n    hooks:\n      - id: end-of-file-fixer\n      - id: trailing-whitespace\n"
			return s.FS.WriteFile(filepath.Join(req.ProjectPath, ".pre-commit-config.yaml"), []byte(cfg), 0o644)
		}})
	}
	if contains(req.Extras, "sentry") {
		steps = append(steps, runner.Step{ID: "sentry", Title: "Adding Sentry config notes", Timeout: 5 * time.Second, Action: func(context.Context) error {
			return s.FS.WriteFile(filepath.Join(req.ProjectPath, "SENTRY.md"), []byte("Set SENTRY_DSN in .env and initialize SDK in app startup.\n"), 0o644)
		}})
	}
	if contains(req.Extras, "docker") {
		steps = append(steps, runner.Step{ID: "docker", Title: "Setting up Docker", Timeout: 5 * time.Second, Action: func(context.Context) error {
			dockerfile := "FROM alpine:3.20\nCMD [\"sh\",\"-c\",\"echo forge scaffold\"]\n"
			compose := "services:\n  app:\n    build: .\n"
			if err := s.FS.WriteFile(filepath.Join(req.ProjectPath, "Dockerfile"), []byte(dockerfile), 0o644); err != nil {
				return err
			}
			return s.FS.WriteFile(filepath.Join(req.ProjectPath, "docker-compose.yml"), []byte(compose), 0o644)
		}})
	}
	if contains(req.Extras, "git") || s.Config.Defaults.GitInit {
		steps = append(steps, runner.CommandStep(
			s.Executor,
			system.Command{Name: "git", Args: []string{"init"}, Dir: req.ProjectPath, Stdout: cmdOut, Stderr: cmdOut},
			"git-init",
			"Git init",
			15*time.Second,
			0,
			"git init",
		))
	}
	return steps, nil
}

func appendLineIfMissing(path, line string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	text := string(content)
	if strings.Contains(text, line) {
		return content, nil
	}
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	text += line + "\n"
	return []byte(text), nil
}

func appendMultiIfMissing(path string, lines []string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	text := string(content)
	for _, line := range lines {
		if strings.Contains(text, line) {
			continue
		}
		if !strings.HasSuffix(text, "\n") {
			text += "\n"
		}
		text += line + "\n"
	}
	return []byte(text), nil
}

func contains(items []string, x string) bool {
	for _, it := range items {
		if strings.EqualFold(strings.TrimSpace(it), x) {
			return true
		}
	}
	return false
}

func ParseExtras(raw string) []string {
	parts := strings.Split(raw, ",")
	seen := map[string]struct{}{}
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		p = strings.ToLower(p)
		if _, ok := seen[p]; ok {
			continue
		}
		seen[p] = struct{}{}
		out = append(out, p)
	}
	sort.Strings(out)
	return out
}
