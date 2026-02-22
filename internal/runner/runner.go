package runner

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type FailurePolicy int

const (
	FailImmediate FailurePolicy = iota
	FailSkippable
)

type Step struct {
	ID         string
	Title      string
	Timeout    time.Duration
	Retry      int
	Policy     FailurePolicy
	DryRunText string
	Action     func(ctx context.Context) error
}

type EventType string

const (
	EventStarted EventType = "started"
	EventDone    EventType = "done"
	EventFailed  EventType = "failed"
	EventSkipped EventType = "skipped"
)

type Event struct {
	Type    EventType
	Index   int
	Total   int
	StepID  string
	Title   string
	Err     error
	Elapsed time.Duration
}

type FailureResolver func(step Step, err error) (decision string)

type Runner struct {
	DryRun       bool
	Observer     func(Event)
	ResolveError FailureResolver
}

func (r Runner) emit(e Event) {
	if r.Observer != nil {
		r.Observer(e)
	}
}

func (r Runner) Run(ctx context.Context, steps []Step) error {
	started := time.Now()
	for idx := 0; idx < len(steps); idx++ {
		step := steps[idx]
		r.emit(Event{Type: EventStarted, Index: idx + 1, Total: len(steps), StepID: step.ID, Title: step.Title})
		if r.DryRun {
			r.emit(Event{Type: EventDone, Index: idx + 1, Total: len(steps), StepID: step.ID, Title: step.Title})
			continue
		}
		attempts := step.Retry + 1
		if attempts < 1 {
			attempts = 1
		}
		var lastErr error
		for attempt := 0; attempt < attempts; attempt++ {
			stepCtx := ctx
			cancel := func() {}
			if step.Timeout > 0 {
				stepCtx, cancel = context.WithTimeout(ctx, step.Timeout)
			}
			err := step.Action(stepCtx)
			cancel()
			if err == nil {
				lastErr = nil
				break
			}
			lastErr = err
			if errors.Is(err, context.DeadlineExceeded) && attempt < attempts-1 {
				continue
			}
		}
		if lastErr != nil {
			r.emit(Event{Type: EventFailed, Index: idx + 1, Total: len(steps), StepID: step.ID, Title: step.Title, Err: lastErr, Elapsed: time.Since(started)})
			decision := "abort"
			if r.ResolveError != nil {
				decision = r.ResolveError(step, lastErr)
			}
			switch decision {
			case "skip":
				if step.Policy != FailSkippable {
					return fmt.Errorf("step %s is not skippable: %w", step.ID, lastErr)
				}
				r.emit(Event{Type: EventSkipped, Index: idx + 1, Total: len(steps), StepID: step.ID, Title: step.Title})
				continue
			case "retry":
				idx--
				continue
			default:
				return fmt.Errorf("step %s failed: %w", step.ID, lastErr)
			}
		}
		r.emit(Event{Type: EventDone, Index: idx + 1, Total: len(steps), StepID: step.ID, Title: step.Title, Elapsed: time.Since(started)})
	}
	return nil
}
