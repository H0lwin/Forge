package cmd

import (
	"fmt"

	"forge/internal/app"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func newDoctorCmd(d *rootDeps) *cobra.Command {
	var noInteractive bool
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Check your development environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd
			_ = args
			res := d.services.Doctor(d.ctx)
			fmt.Fprintln(d.stdout, "Checking your development environment...")
			app.RenderDoctorTable(d.stdout, res.Reports)
			if len(res.Missing) == 0 || noInteractive {
				return nil
			}
			var selected string
			opts := []huh.Option[string]{}
			for _, m := range res.Missing {
				opts = append(opts, huh.NewOption("Install "+m.Tool, m.Tool))
			}
			opts = append(opts, huh.NewOption("Skip for now", "skip"))
			form := huh.NewForm(huh.NewGroup(huh.NewSelect[string]().Title("Missing tools detected").Options(opts...).Value(&selected)))
			if err := form.Run(); err != nil {
				return err
			}
			if selected != "" && selected != "skip" {
				fmt.Fprintf(d.stdout, "Install hint for %s: %s\n", selected, d.services.InstallHint(selected))
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&noInteractive, "no-interactive", false, "disable prompts")
	return cmd
}
