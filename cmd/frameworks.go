package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newFrameworksCmd(d *rootDeps) *cobra.Command {
	_ = d
	return &cobra.Command{
		Use:   "frameworks",
		Short: "List supported frameworks and notable extras",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd
			_ = args
			fmt.Fprintln(d.stdout, "Supported Frameworks")
			fmt.Fprintln(d.stdout, "  Backend:  django, fastapi, flask, express, nestjs")
			fmt.Fprintln(d.stdout, "  Frontend: next, vite")
			fmt.Fprintln(d.stdout, "")
			fmt.Fprintln(d.stdout, "Common Extras: git, docker, ci, pytest, sentry, precommit")
			fmt.Fprintln(d.stdout, "Python Extras: postgres, celery")
			fmt.Fprintln(d.stdout, "Django Extra: drf")
			fmt.Fprintln(d.stdout, "Frontend Extra: tailwind")
			return nil
		},
	}
}
