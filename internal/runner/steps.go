package runner

import (
	"context"
	"io"
	"time"

	"forge/internal/system"
)

func CommandStep(ex system.Executor, cmd system.Command, id, title string, timeout time.Duration, retry int, dryRunText string) Step {
	return Step{
		ID:         id,
		Title:      title,
		Timeout:    timeout,
		Retry:      retry,
		DryRunText: dryRunText,
		Action: func(ctx context.Context) error {
			if cmd.Stdout == nil {
				cmd.Stdout = io.Discard
			}
			if cmd.Stderr == nil {
				cmd.Stderr = io.Discard
			}
			return ex.Run(ctx, cmd)
		},
	}
}
