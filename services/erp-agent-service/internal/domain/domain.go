package domain

import (
	"context"
	"time"
)

type Runner interface {
	RunCycle(ctx context.Context) error
}

type MonitorHandler interface {
	Fetch(ctx context.Context, since time.Time) error
}

type ERPIngestClient interface {
	Post(ctx context.Context, payload any) error
}
