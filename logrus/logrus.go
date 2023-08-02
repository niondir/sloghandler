package logrus

import (
	"context"
	"github.com/sirupsen/logrus"
	"log/slog"
	"slices"
	"strings"
)

var _ slog.Handler = &Handler{}

// Logrus error levels
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel = slog.LevelError + 2
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel = slog.LevelError + 1
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = slog.LevelError
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = slog.LevelWarn
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel = slog.LevelInfo
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = slog.LevelDebug
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel = slog.LevelDebug - 1
)

func SlogLevelToLogrusLevel(level slog.Level) logrus.Level {
	if level <= TraceLevel {
		return logrus.TraceLevel
	} else if level <= DebugLevel {
		return logrus.DebugLevel
	} else if level <= InfoLevel {
		return logrus.InfoLevel
	} else if level <= WarnLevel {
		return logrus.WarnLevel
	} else if level <= ErrorLevel {
		return logrus.ErrorLevel
	} else if level <= FatalLevel {
		return logrus.FatalLevel
	} else {
		return logrus.PanicLevel
	}
}

type Handler struct {
	logger *logrus.Logger
	groups []string
	attrs  []slog.Attr
}

func NewHandler(logger *logrus.Logger) *Handler {
	if logger == nil {
		panic("logger is nil")
	}
	return &Handler{
		logger: logger,
	}
}

func (l *Handler) clone() *Handler {
	return &Handler{
		logger: l.logger,
		groups: slices.Clip(l.groups),
		attrs:  slices.Clip(l.attrs),
	}
}

func (l *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	logrusLevel := l.logger.GetLevel()
	switch logrusLevel {
	case logrus.TraceLevel:
		return true
	case logrus.DebugLevel:
		return level >= slog.LevelDebug
	case logrus.InfoLevel:
		return level >= slog.LevelInfo
	case logrus.WarnLevel:
		return level >= slog.LevelWarn
	case logrus.ErrorLevel:
		return level >= slog.LevelError
	case logrus.FatalLevel:
		return level >= slog.LevelError+1
	case logrus.PanicLevel:
		return level >= slog.LevelError+2
	default:
		return false
	}

}

func (l *Handler) Handle(ctx context.Context, r slog.Record) error {
	log := logrus.NewEntry(l.logger)
	if r.Time.IsZero() {
		log = log.WithTime(r.Time)
	}
	log = log.WithFields(attrsToFields(l.attrs))
	r.Attrs(func(attr slog.Attr) bool {
		if attr.Key == "" {
			return true
		}
		log = log.WithField(attr.Key, attr.Value)
		return true
	})
	log.Logf(SlogLevelToLogrusLevel(r.Level), r.Message)
	return nil
}

func attrsToFields(attrs []slog.Attr) logrus.Fields {
	f := logrus.Fields{}
	for _, a := range attrs {
		if a.Key != "" {
			f[a.Key] = a.Value
		}
	}
	return f
}

func (l *Handler) groupPrefix() string {
	const groupSep = ":"
	if len(l.groups) > 0 {
		return strings.Join(l.groups, groupSep) + groupSep
	}
	return ""
}

func (l *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := l.clone()
	for _, a := range attrs {
		newHandler.attrs = append(newHandler.attrs, slog.Attr{
			Key:   l.groupPrefix() + a.Key,
			Value: a.Value,
		})
	}
	return newHandler
}

func (l *Handler) WithGroup(name string) slog.Handler {
	newHandler := l.clone()
	newHandler.groups = append(newHandler.groups, name)
	return newHandler
}
