package generator

import (
	"context"

	"forge/internal/domain"
	"forge/internal/runner"
)

type Generator interface {
	Name() string
	Category() string
	Steps(ctx context.Context, req domain.GenerateRequest) []runner.Step
	PreCheck(ctx context.Context, req domain.GenerateRequest) error
	PostMessage(req domain.GenerateRequest, result domain.GenerateResult) string
}
