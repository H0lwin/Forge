package cmd

import (
	"fmt"
	"path/filepath"

	"forge/internal/app"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func newAddCmd(d *rootDeps) *cobra.Command {
	var addon string
	var path string
	var noInteractive bool
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a feature to an existing project",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd
			_ = args
			framework := app.DetectFrameworkFromPath(path)
			if framework == "" {
				return fmt.Errorf("could not detect supported project in %s", path)
			}
			fmt.Fprintf(d.stdout, "Detected project: %s\n", framework)
			if addon == "" {
				if noInteractive {
					return fmt.Errorf("--feature is required when --no-interactive is set")
				}
				options := []huh.Option[string]{
					huh.NewOption("auth", "auth"),
					huh.NewOption("celery", "celery"),
					huh.NewOption("cache", "cache"),
					huh.NewOption("email", "email"),
					huh.NewOption("docker", "docker"),
					huh.NewOption("ci", "ci"),
					huh.NewOption("sentry", "sentry"),
					huh.NewOption("pytest", "pytest"),
				}
				if err := huh.NewForm(
					huh.NewGroup(
						huh.NewSelect[string]().
							Title("What would you like to add?").
							Options(options...).
							Value(&addon),
					),
				).Run(); err != nil {
					return err
				}
			}
			rel, content, err := app.AddonFile(addon)
			if err != nil {
				return err
			}
			if err := d.services.FS.WriteFile(filepath.Join(path, rel), []byte(content), 0o644); err != nil {
				return err
			}
			fmt.Fprintf(d.stdout, "Added %s via %s\n", addon, rel)
			return nil
		},
	}
	cmd.Flags().StringVar(&addon, "feature", "", "feature to add: auth|celery|cache|email|docker|ci|sentry|pytest")
	cmd.Flags().StringVar(&addon, "addon", "", "alias for --feature")
	cmd.Flags().StringVar(&path, "path", ".", "project path")
	cmd.Flags().BoolVar(&noInteractive, "no-interactive", false, "disable prompts")
	return cmd
}
