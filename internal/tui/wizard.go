package tui

import (
	"fmt"
	"strings"

	"forge/internal/domain"

	"github.com/charmbracelet/huh"
)

type WizardInput struct {
	DefaultPath          string
	DefaultPythonVersion string
	DefaultEnvManager    string
}

type frameworkOption struct {
	label     string
	framework domain.Framework
	category  domain.ProjectCategory
}

var allFrameworkOptions = []frameworkOption{
	{label: "Django", framework: domain.FrameworkDjango, category: domain.CategoryBackend},
	{label: "FastAPI", framework: domain.FrameworkFastAPI, category: domain.CategoryBackend},
	{label: "Flask", framework: domain.FrameworkFlask, category: domain.CategoryBackend},
	{label: "Express", framework: domain.FrameworkExpress, category: domain.CategoryBackend},
	{label: "NestJS", framework: domain.FrameworkNestJS, category: domain.CategoryBackend},
	{label: "Next.js", framework: domain.FrameworkNext, category: domain.CategoryFrontend},
	{label: "Vite", framework: domain.FrameworkVite, category: domain.CategoryFrontend},
}

func RunWizard(input WizardInput) (domain.GenerateRequest, error) {
	var req domain.GenerateRequest
	var category string
	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(stepTitle(1, 6, "What type of project are you building?")).
				Options(
					huh.NewOption("Backend", string(domain.CategoryBackend)),
					huh.NewOption("Frontend", string(domain.CategoryFrontend)),
				).
				Value(&category),
		),
	).Run(); err != nil {
		return req, err
	}
	req.Category = domain.ProjectCategory(category)

	frameworkOptions := makeFrameworkOptions(req.Category)
	if len(frameworkOptions) == 0 {
		return req, fmt.Errorf("no frameworks available for category %s", req.Category)
	}

	var framework string
	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title(stepTitle(2, 6, "Choose your framework")).Options(frameworkOptions...).Value(&framework),
		),
	).Run(); err != nil {
		return req, err
	}
	req.Framework = domain.Framework(framework)

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title(stepTitle(3, 6, "Project name")).Placeholder("my-awesome-app").Value(&req.Name),
			huh.NewInput().Title("Project path").Value(&req.BasePath).Placeholder(input.DefaultPath),
		),
	).Run(); err != nil {
		return req, err
	}

	if domain.RequiresPythonVersion(req.Framework) {
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title(stepTitle(4, 6, "Python version")).Value(&req.PythonVersion).Placeholder(input.DefaultPythonVersion),
				huh.NewSelect[string]().Title("Environment manager").Options(
					huh.NewOption("venv", "venv"),
					huh.NewOption("poetry", "poetry"),
					huh.NewOption("uv", "uv"),
				).Value(&req.EnvManager),
			),
		).Run(); err != nil {
			return req, err
		}
	}

	extraOpts := extraOptionsForFramework(req.Framework)
	var extras []string
	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().Title(stepTitle(5, 6, "Select extras")).Options(extraOpts...).Value(&extras),
		),
	).Run(); err != nil {
		return req, err
	}
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

	var confirm bool
	review := fmt.Sprintf(
		"Category: %s\nFramework: %s\nName: %s\nPath: %s\nPython: %s\nEnv: %s\nExtras: %s",
		req.Category,
		req.Framework,
		req.Name,
		req.BasePath,
		req.PythonVersion,
		req.EnvManager,
		strings.Join(req.Extras, ", "),
	)
	fmt.Printf("\n%s\n%s\n\n", stepTitle(6, 6, "Review before creation"), review)
	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Create project with this configuration?").Affirmative("Create").Negative("Cancel").Value(&confirm),
		),
	).Run(); err != nil {
		return req, err
	}
	if !confirm {
		return req, fmt.Errorf("project creation canceled")
	}
	return req, nil
}

func makeFrameworkOptions(category domain.ProjectCategory) []huh.Option[string] {
	out := []huh.Option[string]{}
	for _, item := range allFrameworkOptions {
		if item.category != category {
			continue
		}
		out = append(out, huh.NewOption(item.label, string(item.framework)))
	}
	return out
}

func extraOptionsForFramework(f domain.Framework) []huh.Option[string] {
	common := []huh.Option[string]{
		huh.NewOption("git", "git"),
		huh.NewOption("docker", "docker"),
		huh.NewOption("ci", "ci"),
		huh.NewOption("pytest", "pytest"),
		huh.NewOption("sentry", "sentry"),
		huh.NewOption("precommit", "precommit"),
	}
	if f == domain.FrameworkNext || f == domain.FrameworkVite {
		return append(common, huh.NewOption("tailwind", "tailwind"))
	}
	if f == domain.FrameworkDjango {
		return append(common,
			huh.NewOption("drf", "drf"),
			huh.NewOption("postgres", "postgres"),
			huh.NewOption("celery", "celery"),
		)
	}
	if f == domain.FrameworkFastAPI || f == domain.FrameworkFlask {
		return append(common,
			huh.NewOption("postgres", "postgres"),
			huh.NewOption("celery", "celery"),
		)
	}
	return common
}

func stepTitle(step, total int, title string) string {
	return fmt.Sprintf("Step %d/%d - %s", step, total, title)
}
