package app

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"forge/internal/config"
	"forge/internal/domain"
	"forge/internal/generator"
	"forge/internal/generator/backend"
	"forge/internal/generator/frontend"
	"forge/internal/system"
	"forge/internal/templates"
)

type fakeExec struct{}

func (fakeExec) Run(context.Context, system.Command) error { return nil }
func (fakeExec) LookPath(string) (string, error)          { return "ok", nil }

func testServices(t *testing.T) *Services {
	t.Helper()
	eng, err := templates.NewEngine()
	if err != nil {
		t.Fatalf("templates: %v", err)
	}
	fs := system.OSFileSystem{}
	ex := fakeExec{}
	reg := generator.NewRegistry(
		backend.Django(fs, ex),
		backend.FastAPI(fs, ex),
		backend.Flask(fs, ex),
		backend.Express(fs, ex),
		backend.NestJS(fs, ex),
		frontend.Next(fs, ex),
		frontend.Vite(fs, ex),
	)
	cfg := config.Default()
	cfg.Defaults.GitInit = false
	return &Services{FS: fs, Executor: ex, Config: cfg, Templates: eng, Registry: reg}
}

func TestGenerateDryRunAllPhase1Frameworks(t *testing.T) {
	s := testServices(t)
	frameworks := []domain.Framework{
		domain.FrameworkDjango, domain.FrameworkFastAPI, domain.FrameworkFlask,
		domain.FrameworkExpress, domain.FrameworkNestJS, domain.FrameworkNext, domain.FrameworkVite,
	}
	for _, fw := range frameworks {
		req := domain.GenerateRequest{
			Name:          "demo-app",
			Framework:     fw,
			BasePath:      t.TempDir(),
			PythonVersion: "3.11",
			EnvManager:    "venv",
			DryRun:        true,
			NoInteractive: true,
		}
		if _, err := s.Generate(context.Background(), req, io.Discard, nil, nil); err != nil {
			t.Fatalf("framework %s: %v", fw, err)
		}
	}
}

func TestGenerateCreatesRepresentativeProjects(t *testing.T) {
	s := testServices(t)
	cases := []domain.Framework{domain.FrameworkDjango, domain.FrameworkFastAPI, domain.FrameworkNext}
	for _, fw := range cases {
		base := t.TempDir()
		req := domain.GenerateRequest{Name: "demo-app", Framework: fw, BasePath: base, PythonVersion: "3.11", EnvManager: "venv", NoInteractive: true}
		res, err := s.Generate(context.Background(), req, io.Discard, nil, nil)
		if err != nil {
			t.Fatalf("generate %s: %v", fw, err)
		}
		if _, err := os.Stat(filepath.Join(res.ProjectPath, "README.md")); err != nil {
			t.Fatalf("missing README for %s: %v", fw, err)
		}
	}
}
