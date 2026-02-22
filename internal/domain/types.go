package domain

import "time"

type ProjectCategory string

const (
	CategoryFrontend ProjectCategory = "frontend"
	CategoryBackend  ProjectCategory = "backend"
	CategoryMobile   ProjectCategory = "mobile"
)

type Framework string

const (
	FrameworkDjango  Framework = "django"
	FrameworkFastAPI Framework = "fastapi"
	FrameworkFlask   Framework = "flask"
	FrameworkExpress Framework = "express"
	FrameworkNestJS  Framework = "nestjs"
	FrameworkNext    Framework = "next"
	FrameworkVite    Framework = "vite"
)

type GenerateRequest struct {
	Name          string
	Framework     Framework
	Category      ProjectCategory
	BasePath      string
	ProjectPath   string
	PythonVersion string
	EnvManager    string
	Extras        []string
	NoInteractive bool
	DryRun        bool
	Verbose       bool
}

type GenerateResult struct {
	ProjectPath string
	StartedAt   time.Time
	EndedAt     time.Time
	StepsTotal  int
	StepsDone   int
	Skipped     []string
}
