package cmd

import (
	"context"
	"fmt"
	"io"
	"strings"

	"forge/internal/app"
	"forge/internal/cli"
	"forge/internal/domain"

	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
	commit  = "dev"
	date    = "unknown"
)

type rootDeps struct {
	ctx        context.Context
	stdout     io.Writer
	stderr     io.Writer
	verbose    bool
	configPath string
	services   *app.Services
}

func Execute(ctx context.Context, stdout, stderr io.Writer) error {
	cmd := NewRootCommand(ctx, stdout, stderr)
	return cmd.Execute()
}

func NewRootCommand(ctx context.Context, stdout, stderr io.Writer) *cobra.Command {
	d := &rootDeps{ctx: ctx, stdout: stdout, stderr: stderr}
	root := &cobra.Command{
		Use:   "forge",
		Short: "Professional project scaffolding",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(stdout, cli.Banner(version))
			fmt.Fprintln(stdout)
			fmt.Fprintln(stdout, "Available Commands:")
			fmt.Fprintln(stdout, "  - new        Create a new project interactively")
			fmt.Fprintln(stdout, "  - add        Add a feature to existing project")
			fmt.Fprintln(stdout, "  - doctor     Check your development environment")
			fmt.Fprintln(stdout, "  - config     Manage forge settings")
			fmt.Fprintln(stdout, "  - templates  Browse and manage project templates")
			fmt.Fprintln(stdout, "  - version    Show version info")
			fmt.Fprintln(stdout)
			fmt.Fprintln(stdout, "Run 'forge new' to get started")
			_ = cmd
			_ = args
			return nil
		},
		SilenceUsage: true,
	}
	root.PersistentFlags().BoolVar(&d.verbose, "verbose", false, "show raw command output")
	root.PersistentFlags().StringVar(&d.configPath, "config", "", "config path")
	root.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		_ = args
		if cmd.Name() == "version" {
			return nil
		}
		s, err := app.NewServices(stdout, d.verbose, d.configPath)
		if err != nil {
			return err
		}
		d.services = s
		return nil
	}
	root.AddCommand(newNewCmd(d), newDoctorCmd(d), newAddCmd(d), newConfigCmd(d), newTemplatesCmd(d), newVersionCmd(stdout))
	return root
}

func parseFramework(s string) (domain.Framework, error) {
	raw := strings.ToLower(strings.TrimSpace(s))
	switch raw {
	case "nextjs":
		raw = "next"
	}
	f := domain.Framework(raw)
	switch f {
	case domain.FrameworkDjango, domain.FrameworkFastAPI, domain.FrameworkFlask, domain.FrameworkExpress, domain.FrameworkNestJS, domain.FrameworkNext, domain.FrameworkVite:
		return f, nil
	default:
		return "", fmt.Errorf("unsupported framework: %s", s)
	}
}
