package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func newVersionCmd(out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version info",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd
			_ = args
			_, err := fmt.Fprintf(out, "forge %s\ncommit=%s\nbuildDate=%s\n", version, commit, date)
			return err
		},
	}
}
