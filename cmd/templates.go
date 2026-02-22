package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func newTemplatesCmd(d *rootDeps) *cobra.Command {
	return &cobra.Command{
		Use:   "templates",
		Short: "Browse and manage templates",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd
			_ = args
			list := d.services.Templates.List()
			sort.Strings(list)
			fmt.Fprintf(d.stdout, "Embedded templates (%d):\n", len(list))
			for _, t := range list {
				fmt.Fprintf(d.stdout, "  - %s\n", t)
			}
			return nil
		},
	}
}
