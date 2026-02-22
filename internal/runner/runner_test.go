package runner

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunnerRetrySuccess(t *testing.T) {
	attempts := 0
	r := Runner{}
	err := r.Run(context.Background(), []Step{{
		ID: "x", Title: "retry", Retry: 2,
		Action: func(context.Context) error {
			attempts++
			if attempts < 2 {
				return errors.New("boom")
			}
			return nil
		},
	}})
	if err != nil {
		t.Fatalf("run error: %v", err)
	}
	if attempts != 2 {
		t.Fatalf("expected 2 attempts, got %d", attempts)
	}
}

func TestRunnerTimeout(t *testing.T) {
	r := Runner{}
	err := r.Run(context.Background(), []Step{{
		ID: "t", Title: "timeout", Timeout: 10 * time.Millisecond,
		Action: func(ctx context.Context) error {
			<-ctx.Done()
			return ctx.Err()
		},
	}})
	if err == nil {
		t.Fatalf("expected timeout error")
	}
}
