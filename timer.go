package main

import (
	"context"
	"fmt"
	"time"
)

var timerSinceCtxKey = ctxKey("timerSince")

func timerWithSince(ctx context.Context) context.Context {
	return context.WithValue(ctx, timerSinceCtxKey, time.Now())
}

func startTimer(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(timerWithSince(ctx), d)
}

func getTimerSince(ctx context.Context) time.Duration {
	return time.Since(ctx.Value(timerSinceCtxKey).(time.Time))
}

func checkTimeout(ctx context.Context) {
	select {
	case <-ctx.Done():
		d := getTimerSince(ctx)

		if newRelicEnabled {
			sendTimeoutToNewRelic(ctx, d)
		}

		if prometheusEnabled {
			incrementPrometheusErrorsTotal("timeout")
		}

		panic(newError(503, fmt.Sprintf("Timeout after %v", d), "Timeout"))
	default:
		// Go ahead
	}
}
