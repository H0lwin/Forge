package generator

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"forge/internal/domain"
	"forge/internal/runner"
	"forge/internal/system"
)

type Builder struct {
	FS       system.FileSystem
	Executor system.Executor
}

type Spec struct {
	Name      string
	Category  string
	Tools     []string
	Files     map[string]string
	Next      []string
	Bootstrap []system.Command
}

type genericGenerator struct {
	b    Builder
	spec Spec
}

func NewGenericBuilder(fs system.FileSystem, ex system.Executor) Builder {
	return Builder{FS: fs, Executor: ex}
}

func (b Builder) New(spec Spec) Generator {
	return &genericGenerator{b: b, spec: spec}
}

func (g *genericGenerator) Name() string { return g.spec.Name }
func (g *genericGenerator) Category() string { return g.spec.Category }

func (g *genericGenerator) PreCheck(ctx context.Context, req domain.GenerateRequest) error {
	_ = req
	for _, t := range g.spec.Tools {
		alternatives := strings.Split(t, "|")
		ok := false
		for _, candidate := range alternatives {
			candidate = strings.TrimSpace(candidate)
			if candidate == "" {
				continue
			}
			if _, err := g.b.Executor.LookPath(candidate); err == nil {
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf("%w: %s", domain.ErrToolNotFound, t)
		}
	}
	return nil
}

func (g *genericGenerator) Steps(ctx context.Context, req domain.GenerateRequest) []runner.Step {
	_ = ctx
	steps := []runner.Step{
		{
			ID:      "mkdir",
			Title:   "Creating project directory",
			Timeout: 15 * time.Second,
			Action: func(_ context.Context) error {
				return g.b.FS.MkdirAll(req.ProjectPath, 0o755)
			},
		},
	}
	paths := make([]string, 0, len(g.spec.Files))
	for rel := range g.spec.Files {
		paths = append(paths, rel)
	}
	sort.Strings(paths)
	for _, rel := range paths {
		content := g.spec.Files[rel]
		path := filepath.Join(req.ProjectPath, rel)
		c := content
		steps = append(steps, runner.Step{
			ID:      "file_" + rel,
			Title:   "Writing " + rel,
			Timeout: 10 * time.Second,
			Action: func(_ context.Context) error {
				return g.b.FS.WriteFile(path, []byte(c), 0o644)
			},
		})
	}
	for i, c := range g.spec.Bootstrap {
		cmd := c
		cmd.Dir = req.ProjectPath
		steps = append(steps, runner.CommandStep(g.b.Executor, cmd, fmt.Sprintf("bootstrap_%d", i+1), "Running bootstrap command", 2*time.Minute, 0, cmd.Name))
	}
	return steps
}

func (g *genericGenerator) PostMessage(req domain.GenerateRequest, _ domain.GenerateResult) string {
	return fmt.Sprintf("Project created at %s\nNext: %v", req.ProjectPath, g.spec.Next)
}
