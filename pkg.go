package log

import (
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var defaultLogger atomic.Pointer[Logger]
var defaultOnce sync.Once

func Default() *Logger {
	defaultOnce.Do(func() {
		defaultLogger.CompareAndSwap(nil, New(os.Stdout, WithPrefix("main"), WithTimeFormat(time.TimeOnly)))
	})
	return defaultLogger.Load()
}

func SetDefault(logger *Logger) {
	defaultLogger.Store(logger)
}

func SetLevel(level Level) {
	Default().SetLevel(level)
}

// Log logs a message with the given level.
func Log(level Level, msg any, keyvals ...any) {
	Default().Log(level, msg, keyvals...)
}

// Debug logs a debug message with the given key-value pairs.
func Debug(msg any, keyvals ...any) {
	Default().Log(DebugLevel, msg, keyvals...)
}

// Hint logs a hint message with the given key-value pairs.
func Hint(msg any, keyvals ...any) {
	Default().Log(HintLevel, msg, keyvals...)
}

// Info logs an info message with the given key-value pairs.
func Info(msg any, keyvals ...any) {
	Default().Log(InfoLevel, msg, keyvals...)
}

// Warn logs a warning message with the given key-value pairs.
func Warn(msg any, keyvals ...any) {
	Default().Log(WarnLevel, msg, keyvals...)
}

// Cart logs a cart message with the given key-value pairs.
func Cart(msg any, keyvals ...any) {
	Default().Log(CartLevel, msg, keyvals...)
}

// Miss logs a miss message with the given key-value pairs.
func Miss(msg any, keyvals ...any) {
	Default().Log(MissLevel, msg, keyvals...)
}

// Error logs an error message with the given key-value pairs.
func Error(msg any, keyvals ...any) {
	Default().Log(ErrorLevel, msg, keyvals...)
}

// Done logs a done message with the given key-value pairs.
func Done(msg any, keyvals ...any) {
	Default().Log(DoneLevel, msg, keyvals...)
}

// Fatal logs a fatal message with the given key-value pairs and exits the program.
func Fatal(msg any, keyvals ...any) {
	Default().Log(FatalLevel, msg, keyvals...)
	os.Exit(1)
}

// Print logs a message without a level and with the given key-value pairs.
func Print(msg any, keyvals ...any) {
	Default().Log(noLevel, msg, keyvals...)
}

// Logf logs a message with formatting and level.
func Logf(level Level, format string, args ...any) {
	Default().Logf(level, format, args...)
}

// Debugf prints a debug message with formatting.
func Debugf(format string, args ...any) {
	Default().Logf(DebugLevel, format, args...)
}

// Hintf prints a hint message with formatting.
func Hintf(format string, args ...any) {
	Default().Logf(HintLevel, format, args...)
}

// Infof prints an info message with formatting.
func Infof(format string, args ...any) {
	Default().Logf(InfoLevel, format, args...)
}

// Warnf prints a warning message with formatting.
func Warnf(format string, args ...any) {
	Default().Logf(WarnLevel, format, args...)
}

// Cartf prints a cart message with formatting.
func Cartf(format string, args ...any) {
	Default().Logf(CartLevel, format, args...)
}

// Missf prints a miss message with formatting.
func Missf(format string, args ...any) {
	Default().Logf(MissLevel, format, args...)
}

// Errorf prints an error message with formatting.
func Errorf(format string, args ...any) {
	Default().Logf(ErrorLevel, format, args...)
}

// Donef prints a done message with formatting.
func Donef(format string, args ...any) {
	Default().Logf(DoneLevel, format, args...)
}

// Fatalf prints a fatal message with formatting and exits the program.
func Fatalf(format string, args ...any) {
	Default().Logf(FatalLevel, format, args...)
	os.Exit(1)
}
