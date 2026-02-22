package cmd

import (
	"fmt"
	"os"

	"forge/internal/config"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newConfigCmd(d *rootDeps) *cobra.Command {
	var reset bool
	var exportPath string
	var importPath string
	var edit bool
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage forge settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd
			_ = args
			actionCount := 0
			if reset {
				actionCount++
			}
			if exportPath != "" {
				actionCount++
			}
			if importPath != "" {
				actionCount++
			}
			if edit {
				actionCount++
			}
			if actionCount > 1 {
				return fmt.Errorf("use only one action at a time: --reset, --edit, --import, or --export")
			}
			cfg := d.services.Config
			if reset {
				cfg = config.Default()
				if err := config.Save(d.configPath, cfg); err != nil {
					return err
				}
				fmt.Fprintln(d.stdout, "Config reset to defaults")
				return nil
			}
			if importPath != "" {
				b, err := os.ReadFile(importPath)
				if err != nil {
					return err
				}
				if err := yaml.Unmarshal(b, &cfg); err != nil {
					return err
				}
				if err := config.Save(d.configPath, cfg); err != nil {
					return err
				}
				fmt.Fprintf(d.stdout, "Imported config from %s\n", importPath)
				return nil
			}
			if exportPath != "" {
				b, err := yaml.Marshal(cfg)
				if err != nil {
					return err
				}
				if err := os.WriteFile(exportPath, b, 0o644); err != nil {
					return err
				}
				fmt.Fprintf(d.stdout, "Exported config to %s\n", exportPath)
				return nil
			}
			if edit {
				gitInit := cfg.Defaults.GitInit
				openEditor := cfg.Defaults.OpenEditor
				if err := huh.NewForm(
					huh.NewGroup(
						huh.NewInput().Title("User name").Value(&cfg.User.Name),
						huh.NewInput().Title("User email").Value(&cfg.User.Email),
						huh.NewInput().Title("GitHub username").Value(&cfg.User.Github),
						huh.NewInput().Title("Default editor").Value(&cfg.Defaults.Editor),
						huh.NewInput().Title("Python default version").Value(&cfg.Python.DefaultVersion),
						huh.NewSelect[string]().
							Title("Python env manager").
							Options(
								huh.NewOption("venv", "venv"),
								huh.NewOption("poetry", "poetry"),
								huh.NewOption("uv", "uv"),
							).
							Value(&cfg.Python.EnvManager),
						huh.NewInput().Title("Node package manager").Value(&cfg.Node.PackageManager),
						huh.NewConfirm().Title("Git init by default").Value(&gitInit),
						huh.NewConfirm().Title("Open editor after generation").Value(&openEditor),
					),
				).Run(); err != nil {
					return err
				}
				cfg.Defaults.GitInit = gitInit
				cfg.Defaults.OpenEditor = openEditor
				if err := config.Save(d.configPath, cfg); err != nil {
					return err
				}
				fmt.Fprintln(d.stdout, "Config updated")
				return nil
			}
			b, err := yaml.Marshal(cfg)
			if err != nil {
				return err
			}
			fmt.Fprintln(d.stdout, string(b))
			return nil
		},
	}
	cmd.Flags().BoolVar(&reset, "reset", false, "reset to defaults")
	cmd.Flags().BoolVar(&edit, "edit", false, "edit settings interactively")
	cmd.Flags().StringVar(&exportPath, "export", "", "export config file")
	cmd.Flags().StringVar(&importPath, "import", "", "import config file")
	return cmd
}
