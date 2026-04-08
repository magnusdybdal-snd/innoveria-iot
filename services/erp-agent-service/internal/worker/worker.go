// Package worker runs the polling loop for sync cycles.
package worker

import (
	"context"
	"innoveria-iot/erp-agent-service/internal/domain"
	"log/slog"
	"time"
)

// Worker schedules and executes runner cycles with retry backoff.
type Worker struct {
	Interval     time.Duration
	CycleTimeout time.Duration
	MaxBackoff   time.Duration

	Runner domain.Runner
}

// New constructs a worker with interval, timeout, and backoff settings.
func New(interval, cycleTimeout, maxBackoff time.Duration, runner domain.Runner) *Worker {
	return &Worker{
		Interval:     interval,
		CycleTimeout: cycleTimeout,
		MaxBackoff:   maxBackoff,
		Runner:       runner,
	}
}

// Start starts the polling loop until the context is cancelled.
func (w *Worker) Start(ctx context.Context) {
	backoff := time.Second
	nextDelay := time.Duration(0) // run immediately on startup

	for {
		timer := time.NewTimer(nextDelay)
		select {
		case <-ctx.Done(): // Handles shutdown
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			slog.Info("worker loop stopped")
			return
		case <-timer.C: // Handle one cycle
		}

		cycleCtx, cancel := context.WithTimeout(ctx, w.CycleTimeout)
		err := w.Runner.RunCycle(cycleCtx)
		cancel()

		if err != nil {
			slog.Error("worker loop failed", "err", err, "backoff", backoff)
			nextDelay = backoff
			backoff *= 2
			if backoff > w.MaxBackoff {
				backoff = w.MaxBackoff
			}
			continue
		}

		slog.Info("woker loop succeeded")
		backoff = time.Second
		nextDelay = w.Interval
	}
}
