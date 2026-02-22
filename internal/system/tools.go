package system

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var semverRe = regexp.MustCompile(`(\d+\.\d+(?:\.\d+)?)`)

type ToolStatus string

const (
	StatusInstalled ToolStatus = "Installed"
	StatusRunning   ToolStatus = "Running"
	StatusMissing   ToolStatus = "Missing"
)

type ToolReport struct {
	Tool    string
	Status  ToolStatus
	Version string
	Note    string
}

func ParseVersion(text string) string {
	m := semverRe.FindStringSubmatch(text)
	if len(m) < 2 {
		return "-"
	}
	return m[1]
}

func DetectTool(ctx context.Context, ex Executor, name string, versionArgs []string) ToolReport {
	_, err := ex.LookPath(name)
	if err != nil {
		return ToolReport{Tool: name, Status: StatusMissing, Version: "-"}
	}
	buf := &strings.Builder{}
	cmd := Command{Name: name, Args: versionArgs, Stdout: io.MultiWriter(buf), Stderr: io.MultiWriter(buf)}
	if runErr := ex.Run(ctx, cmd); runErr != nil {
		return ToolReport{Tool: name, Status: StatusInstalled, Version: "-", Note: fmt.Sprintf("version check failed: %v", runErr)}
	}
	return ToolReport{Tool: name, Status: StatusInstalled, Version: ParseVersion(buf.String())}
}
