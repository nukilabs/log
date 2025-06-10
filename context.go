package log

import "context"

type contextkey string

const ContextKey contextkey = "log"

func WithContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, ContextKey, logger)
}

func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(ContextKey).(*Logger)
	if !ok {
		return nil
	}
	return logger
}
