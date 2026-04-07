// package worker provides the polling loop to fetch data from monitor erp
package worker

import (
	"context"
	"innoveria-iot/erp-agent-service/internal/domain"
	"log/slog"
	"time"
)

type Worker struct {
	Interval     time.Duration
	CycleTimeout time.Duration
	MaxBackoff   time.Duration

	Runner domain.Runner
}

func NewWorker(interval, cycleTimeout, maxBackoff time.Duration, runner domain.Runner) *Worker {
	return &Worker{
		Interval:     interval,
		CycleTimeout: cycleTimeout,
		MaxBackoff:   maxBackoff,
		Runner:       runner,
	}
}

func (w *Worker) Start(ctx context.Context) {
	ticker := time.NewTicker(w.Interval)
	defer ticker.Stop()

	backoff := time.Second
	running := false // handles duplicated cycles

	for {
		select {
		case <-ctx.Done(): // Handles shutdown
			slog.Info("worker loop stopped")
			return
		case <-ticker.C: // Handle one cycle
			if running {
				slog.Warn("previous worker loop still running")
				continue
			}
			running = true
			func() {
				defer func() { running = false }()

				cycleCtx, cancel := context.WithTimeout(ctx, w.CycleTimeout)
				defer cancel()

				err := w.Runner.RunCycle(cycleCtx)
				if err != nil {
					slog.Error("worker loop failed", "err", err, "backoff", backoff)
					select {
					case <-ctx.Done():
						return
					case <-time.After(backoff):
					}
					backoff *= 2
					if backoff > w.MaxBackoff {
						backoff = w.MaxBackoff
					}
					return
				}
				backoff = time.Second
				slog.Info("woker loop succeeded")
			}()
		}
	}
}
