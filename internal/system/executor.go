package system

import (
	"context"
	"io"
	"os/exec"
	"runtime"
	"strings"
)

type Executor interface {
	Run(ctx context.Context, cmd Command) error
	LookPath(file string) (string, error)
}

type Command struct {
	Name    string
	Args    []string
	Dir     string
	Stdout  io.Writer
	Stderr  io.Writer
	Verbose bool
}

type OSExecutor struct{}

func (e OSExecutor) Run(ctx context.Context, cmd Command) error {
	name := cmd.Name
	args := cmd.Args
	if runtime.GOOS == "windows" && (strings.Contains(cmd.Name, " ") || strings.ContainsAny(cmd.Name, "|&><")) {
		name = "cmd"
		args = []string{"/C", cmd.Name}
	}
	c := exec.CommandContext(ctx, name, args...)
	c.Dir = cmd.Dir
	c.Stdout = cmd.Stdout
	c.Stderr = cmd.Stderr
	return c.Run()
}

func (e OSExecutor) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}
