package customContext

import (
	"context"
	"time"
)

var globalCtx context.Context

func New() {
	globalCtx = context.Background()
}

func WithCancel() (context.Context, context.CancelFunc) {
	return context.WithCancel(globalCtx)
}

func WithDeadline(deadline time.Time) (context.Context, context.CancelFunc) {
	return context.WithDeadline(globalCtx, deadline)
}

func WithTimeout(t time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(globalCtx, t)
}
