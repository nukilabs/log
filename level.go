package log

import (
	"math"
)

// Level is a logging level.
type Level int

const (
	// DebugLevel is the debug level.
	DebugLevel Level = -4
	// HintLevel is the hint level.
	HintLevel Level = 0
	// InfoLevel is the info level.
	InfoLevel Level = 4
	// WarnLevel is the warn level.
	WarnLevel Level = 8
	// CartLevel is the cart level.
	CartLevel Level = 12
	// MissLevel is the miss level.
	MissLevel Level = 16
	// ErrorLevel is the error level.
	ErrorLevel Level = 20
	// DoneLevel is the done level.
	DoneLevel Level = 24
	// FatalLevel is the fatal level.
	FatalLevel Level = 28
	// noLevel is used with log.Print.
	noLevel Level = math.MaxInt
)

// String returns the string representation of the level.
func (l Level) String() string {
	switch l { //nolint:exhaustive
	case DebugLevel:
		return "debug"
	case HintLevel:
		return "hint"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case CartLevel:
		return "cart"
	case MissLevel:
		return "miss"
	case ErrorLevel:
		return "error"
	case DoneLevel:
		return "done"
	case FatalLevel:
		return "fatal"
	default:
		return ""
	}
}
