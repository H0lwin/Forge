package cmd

import (
	"errors"
	"fmt"
	"os"

	"forge/internal/app"
	"forge/internal/domain"
	"forge/internal/runner"
	"forge/internal/tui"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func newNewCmd(d *rootDeps) *cobra.Command {
	var req domain.GenerateRequest
	var frameworkRaw string
	var extras string
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new project",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = args
			if d.services == nil {
				return fmt.Errorf("services not initialized")
			}
			req.Verbose = d.verbose
			if req.NoInteractive {
				req.Extras = app.ParseExtras(extras)
				if frameworkRaw != "" {
					f, err := parseFramework(frameworkRaw)
					if err != nil {
						return err
					}
					req.Framework = f
				}
				missing := domain.MissingRequiredFlags(req)
				if len(missing) > 0 {
					return fmt.Errorf("%w: %v\nexample: forge new --name my-app --framework django --path . --python-version 3.11 --env-manager venv --no-interactive", domain.ErrMissingRequiredFlag, missing)
				}
			} else {
				home, _ := os.UserHomeDir()
				wizReq, err := tui.RunWizard(tui.WizardInput{DefaultPath: home, DefaultPythonVersion: d.services.Config.Python.DefaultVersion, DefaultEnvManager: d.services.Config.Python.EnvManager})
				if err != nil {
					return err
				}
				req = wizReq
				if req.Framework != "" {
					f, err := parseFramework(string(req.Framework))
					if err != nil {
						return err
					}
					req.Framework = f
				}
			}
			observe, stop := tui.NewProgressObserver(d.stdout)
			defer stop()
			_, err := d.services.Generate(d.ctx, req, d.stdout, func(step runner.Step, err error) string {
				if req.NoInteractive {
					return "abort"
				}
				var choice string
				form := huh.NewForm(huh.NewGroup(huh.NewSelect[string]().Title(fmt.Sprintf("%s failed: %v", step.Title, err)).Options(
					huh.NewOption("Skip this step and continue", "skip"),
					huh.NewOption("Retry", "retry"),
					huh.NewOption("Abort setup", "abort"),
				).Value(&choice)))
				if runErr := form.Run(); runErr != nil {
					return "abort"
				}
				if choice == "" {
					choice = "abort"
				}
				return choice
			}, observe)
			if err != nil {
				if errors.Is(err, domain.ErrToolNotFound) && !req.NoInteractive {
					fmt.Fprintln(d.stderr, "A required tool is missing. Run `forge doctor`.")
				}
				return err
			}
			_ = cmd
			return nil
		},
	}
	cmd.Flags().StringVar(&req.Name, "name", "", "project name")
	cmd.Flags().StringVar(&frameworkRaw, "framework", "", "framework")
	cmd.Flags().StringVar(&req.BasePath, "path", "", "base path")
	cmd.Flags().StringVar(&req.PythonVersion, "python-version", "", "python version")
	cmd.Flags().StringVar(&req.EnvManager, "env-manager", "", "python env manager")
	cmd.Flags().StringVar(&extras, "extras", "", "comma-separated extras")
	cmd.Flags().BoolVar(&req.NoInteractive, "no-interactive", false, "disable prompts")
	cmd.Flags().BoolVar(&req.DryRun, "dry-run", false, "show planned actions only")
	return cmd
}
