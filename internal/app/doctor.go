package app

import (
	"context"
	"fmt"
	"io"
	"runtime"
	"strings"

	"forge/internal/system"
)

type DoctorResult struct {
	Reports []system.ToolReport
	Missing []system.ToolReport
}

func (s *Services) Doctor(ctx context.Context) DoctorResult {
	spec := []struct {
		name string
		args []string
	}{
		{"git", []string{"--version"}},
		{"node", []string{"--version"}},
		{"pnpm", []string{"--version"}},
		{"flutter", []string{"--version"}},
		{"uv", []string{"--version"}},
		{"poetry", []string{"--version"}},
	}
	reports := make([]system.ToolReport, 0, len(spec)+2)
	missing := []system.ToolReport{}

	python := system.DetectTool(ctx, s.Executor, "python3", []string{"--version"})
	if python.Status == system.StatusMissing {
		python = system.DetectTool(ctx, s.Executor, "python", []string{"--version"})
		if python.Status != system.StatusMissing {
			python.Tool = "python3"
			python.Note = "resolved via python"
		}
	}
	reports = append(reports, python)
	if python.Status == system.StatusMissing {
		missing = append(missing, python)
	}

	for _, item := range spec {
		r := system.DetectTool(ctx, s.Executor, item.name, item.args)
		reports = append(reports, r)
		if r.Status == system.StatusMissing {
			missing = append(missing, r)
		}
	}

	docker := system.DetectTool(ctx, s.Executor, "docker", []string{"--version"})
	if docker.Status != system.StatusMissing {
		info := system.DetectTool(ctx, s.Executor, "docker", []string{"info", "--format", "{{.ServerVersion}}"})
		if info.Status != system.StatusMissing && info.Version != "-" {
			docker.Status = system.StatusRunning
		}
	}
	reports = append(reports, docker)
	if docker.Status == system.StatusMissing {
		missing = append(missing, docker)
	}

	return DoctorResult{Reports: reports, Missing: missing}
}

func (s *Services) InstallHint(tool string) string {
	osName := runtime.GOOS
	if osName == "windows" {
		switch tool {
		case "uv":
			return "powershell -ExecutionPolicy ByPass -c \"irm https://astral.sh/uv/install.ps1 | iex\""
		case "flutter":
			return "choco install flutter"
		default:
			return "Install manually from official docs"
		}
	}
	switch tool {
	case "uv":
		return "curl -LsSf https://astral.sh/uv/install.sh | sh"
	case "flutter":
		return "See flutter.dev/docs/get-started/install"
	default:
		return "Install manually from official docs"
	}
}

func RenderDoctorTable(w io.Writer, reports []system.ToolReport) {
	fmt.Fprintln(w, "Tool\tStatus\tVersion\tNote")
	for _, r := range reports {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", r.Tool, r.Status, r.Version, strings.TrimSpace(r.Note))
	}
}
