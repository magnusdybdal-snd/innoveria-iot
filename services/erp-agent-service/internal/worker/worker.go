// package worker provides the polling loop to fetch data from monitor erp
package worker

import (
	"context"
	"log/slog"
	"time"
)

type Worker struct {
	Interval     time.Duration
	CycleTimeout time.Duration
	MaxBackoff   time.Duration
}

func (w *Worker) PollingLoop(ctx context.Context) {
	ticker := time.NewTicker(w.Interval)
	defer ticker.Stop()
	running := false

	for {
		select {
		case <-ctx.Done():
			slog.Info("polling loop stopped")
			return
		case <-ticker.C:
			if running {
				slog.Warn("previous sync cycle still running")
				continue
			}
			running = true
			func() {
				defer func() { running = false }()

				//cycleCtx, cancel := context.WithTimeout(ctx, w.CycleTimeout)
			}()
		}
	}
}
