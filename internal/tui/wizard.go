package tui

import (
	"strings"

	"forge/internal/domain"

	"github.com/charmbracelet/huh"
)

type WizardInput struct {
	DefaultPath          string
	DefaultPythonVersion string
	DefaultEnvManager    string
}

func RunWizard(input WizardInput) (domain.GenerateRequest, error) {
	var req domain.GenerateRequest
	var framework string
	var extras []string
	frameworks := []huh.Option[string]{
		huh.NewOption("Django", "django"),
		huh.NewOption("FastAPI", "fastapi"),
		huh.NewOption("Flask", "flask"),
		huh.NewOption("Express", "express"),
		huh.NewOption("NestJS", "nestjs"),
		huh.NewOption("Next.js", "next"),
		huh.NewOption("Vite", "vite"),
	}
	extraOpts := []huh.Option[string]{
		huh.NewOption("git", "git"),
		huh.NewOption("docker", "docker"),
		huh.NewOption("ci", "ci"),
		huh.NewOption("drf", "drf"),
		huh.NewOption("postgres", "postgres"),
		huh.NewOption("pytest", "pytest"),
		huh.NewOption("sentry", "sentry"),
		huh.NewOption("precommit", "precommit"),
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Project name").Value(&req.Name),
			huh.NewSelect[string]().Title("Framework").Options(frameworks...).Value(&framework),
			huh.NewInput().Title("Project path").Value(&req.BasePath).Placeholder(input.DefaultPath),
			huh.NewInput().Title("Python version").Value(&req.PythonVersion).Placeholder(input.DefaultPythonVersion),
			huh.NewSelect[string]().Title("Env manager").Options(
				huh.NewOption("venv", "venv"),
				huh.NewOption("poetry", "poetry"),
				huh.NewOption("uv", "uv"),
			).Value(&req.EnvManager),
			huh.NewMultiSelect[string]().Title("Extras").Options(extraOpts...).Value(&extras),
		),
	)
	if err := form.Run(); err != nil {
		return req, err
	}
	req.Framework = domain.Framework(framework)
	req.Extras = extras
	if strings.TrimSpace(req.BasePath) == "" {
		req.BasePath = input.DefaultPath
	}
	if strings.TrimSpace(req.PythonVersion) == "" {
		req.PythonVersion = input.DefaultPythonVersion
	}
	if strings.TrimSpace(req.EnvManager) == "" {
		req.EnvManager = input.DefaultEnvManager
	}
	req.Category = detectCategory(req.Framework)
	return req, nil
}

func detectCategory(framework domain.Framework) domain.ProjectCategory {
	switch framework {
	case domain.FrameworkNext, domain.FrameworkVite:
		return domain.CategoryFrontend
	default:
		return domain.CategoryBackend
	}
}
