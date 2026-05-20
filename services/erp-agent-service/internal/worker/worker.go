// Package worker runs the polling loop for sync cycles.
package worker

import (
	"context"
	"errors"
	"innoveria-iot/erp-agent-service/internal/domain"
	"log/slog"
	"time"
)

// Worker schedules and executes runner cycles with retry backoff.
type Worker struct {
	interval     time.Duration
	cycleTimeout time.Duration
	maxBackoff   time.Duration

	runner domain.Runner
}

// New constructs a worker with interval, timeout, and backoff settings.
func New(interval, cycleTimeout, maxBackoff time.Duration, runner domain.Runner) *Worker {
	return &Worker{
		interval:     interval,
		cycleTimeout: cycleTimeout,
		maxBackoff:   maxBackoff,
		runner:       runner,
	}
}

// Start starts the polling loop until the context is cancelled or a fatal error occurs.
func (w *Worker) Start(ctx context.Context) error {
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
			return nil
		case <-timer.C: // Handle one cycle
		}

		cycleCtx, cancel := context.WithTimeout(ctx, w.cycleTimeout)
		err := w.runner.RunCycle(cycleCtx)
		cancel()

		if err != nil {
			if errors.Is(err, domain.ErrERPUnauthorized) {
				slog.Error("worker loop stopped due to unauthorized ERP token", "err", err)
				return err
			}
			slog.Error("worker loop failed", "err", err, "backoff", backoff)
			nextDelay = backoff
			backoff *= 2
			if backoff > w.maxBackoff {
				backoff = w.maxBackoff
			}
			continue
		}

		slog.Info("woker loop succeeded")
		backoff = time.Second
		nextDelay = w.interval
	}
}
