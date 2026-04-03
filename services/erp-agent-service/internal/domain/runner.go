package domain

import "context"

type Runner interface {
	RunCycle(ctx context.Context) error
}
