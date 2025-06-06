package log

import (
	"bytes"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	// TimestampKey is the key for the timestamp.
	TimestampKey = "time"
	// MessageKey is the key for the message.
	MessageKey = "msg"
	// IndexKey is the key for the index.
	IndexKey = "index"
	// LevelKey is the key for the level.
	LevelKey = "level"
	// PrefixKey is the key for the prefix.
	PrefixKey = "prefix"
)

type Option func(*Logger)

// WithIndex sets the index for the logger.
func WithIndex(idx int) Option {
	return func(l *Logger) {
		l.idx = idx
	}
}

// WithLevel sets the logging level.
func WithLevel(level Level) Option {
	return func(l *Logger) {
		l.level = int64(level)
	}
}

// WithPrefix sets the prefix for the logger.
func WithPrefix(prefix string) Option {
	return func(l *Logger) {
		l.prefix = prefix
	}
}

// WithTimeFormat sets the time format for the logger.
func WithTimeFormat(format string) Option {
	return func(l *Logger) {
		l.timeFormat = format
	}
}

// WithStyles sets the styles for the logger.
func WithStyles(styles *Styles) Option {
	return func(l *Logger) {
		l.styles = styles
	}
}

// Logger is a structured logger.
type Logger struct {
	w  io.Writer
	b  bytes.Buffer
	mu *sync.RWMutex
	re *lipgloss.Renderer

	idx        int
	level      int64
	prefix     string
	timeFormat string
	styles     *Styles
}

var registry sync.Map

// New creates a new Logger with the given options.
func New(w io.Writer, opts ...Option) *Logger {
	l := &Logger{
		w:          w,
		mu:         &sync.RWMutex{},
		level:      int64(HintLevel),
		idx:        -1,
		prefix:     "",
		timeFormat: time.TimeOnly,
		styles:     DefaultStyles(),
	}

	if v, ok := registry.Load(w); ok {
		l.re = v.(*lipgloss.Renderer)
	} else {
		l.re = lipgloss.NewRenderer(w, termenv.WithColorCache(true))
		registry.Store(w, l.re)
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

// Log logs a message with the given level and key-value pairs.
func (l *Logger) Log(level Level, msg any, keyvals ...any) {
	var kvs []any
	if atomic.LoadInt64(&l.level) > int64(level) {
		return
	}

	kvs = append(kvs, TimestampKey, time.Now())
	if l.idx != -1 {
		kvs = append(kvs, IndexKey, l.idx)
	}
	if _, ok := l.styles.Levels[level]; ok {
		kvs = append(kvs, LevelKey, level)
	}
	if l.prefix != "" {
		kvs = append(kvs, PrefixKey, l.prefix)
	}
	if msg != nil {
		kvs = append(kvs, MessageKey, msg)
	}
	kvs = append(kvs, keyvals...)
	if len(kvs)%2 != 0 {
		kvs = append(kvs, nil)
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	l.formatter(kvs...)
	l.b.WriteTo(l.w)
}

// Debug logs a debug message with the given key-value pairs.
func (l *Logger) Debug(msg any, keyvals ...any) {
	l.Log(DebugLevel, msg, keyvals...)
}

// Hint logs a hint message with the given key-value pairs.
func (l *Logger) Hint(msg any, keyvals ...any) {
	l.Log(HintLevel, msg, keyvals...)
}

// Info logs an info message with the given key-value pairs.
func (l *Logger) Info(msg any, keyvals ...any) {
	l.Log(InfoLevel, msg, keyvals...)
}

// Warn logs a warning message with the given key-value pairs.
func (l *Logger) Warn(msg any, keyvals ...any) {
	l.Log(WarnLevel, msg, keyvals...)
}

// Cart logs a cart message with the given key-value pairs.
func (l *Logger) Cart(msg any, keyvals ...any) {
	l.Log(CartLevel, msg, keyvals...)
}

// Miss logs a miss message with the given key-value pairs.
func (l *Logger) Miss(msg any, keyvals ...any) {
	l.Log(MissLevel, msg, keyvals...)
}

// Error logs an error message with the given key-value pairs.
func (l *Logger) Error(msg any, keyvals ...any) {
	l.Log(ErrorLevel, msg, keyvals...)
}

// Done logs a done message with the given key-value pairs.
func (l *Logger) Done(msg any, keyvals ...any) {
	l.Log(DoneLevel, msg, keyvals...)
}

// Fatal logs a fatal message with the given key-value pairs and exits the program.
func (l *Logger) Fatal(msg any, keyvals ...any) {
	l.Log(FatalLevel, msg, keyvals...)
	os.Exit(1)
}

// Print logs a message without a level and with the given key-value pairs.
func (l *Logger) Print(msg any, keyvals ...any) {
	l.Log(noLevel, msg, keyvals...)
}
