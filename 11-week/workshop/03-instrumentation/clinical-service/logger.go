package clinicalobs

import (
	"context"
	"io"
	"log/slog"
	"time"
)

// NewLogger writes the CP2 log contract: snake_case JSON with timestamp, level,
// service, version, environment, event and message.
func NewLogger(w io.Writer, service, version, env string) *slog.Logger {
	h := slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				return slog.String("timestamp", a.Value.Time().UTC().Format("2006-01-02T15:04:05.000Z"))
			case slog.MessageKey:
				a.Key = "message"
			}
			return a
		},
	})
	return slog.New(traceHandler{h}).With("service", service, "version", version, "environment", env)
}

// traceHandler injects trace_id, span_id and correlation_id from the context.
type traceHandler struct{ slog.Handler }

func (h traceHandler) Handle(ctx context.Context, r slog.Record) error {
	if tc, ok := FromContext(ctx); ok {
		r.AddAttrs(slog.String("trace_id", tc.TraceID), slog.String("span_id", tc.SpanID))
		if id, ok := ctx.Value(corrKey{}).(string); ok {
			r.AddAttrs(slog.String("correlation_id", id))
		} else {
			r.AddAttrs(slog.String("correlation_id", tc.TraceID))
		}
	}
	return h.Handler.Handle(ctx, r)
}

func (h traceHandler) WithAttrs(as []slog.Attr) slog.Handler {
	return traceHandler{h.Handler.WithAttrs(as)}
}
func (h traceHandler) WithGroup(n string) slog.Handler { return traceHandler{h.Handler.WithGroup(n)} }

type corrKey struct{}

// LogEvent writes one contract line; duration_ms is always present.
func LogEvent(ctx context.Context, l *slog.Logger, lvl slog.Level, event, msg string, d time.Duration, attrs ...any) {
	attrs = append([]any{"event", event, "duration_ms", d.Milliseconds()}, attrs...)
	l.Log(ctx, lvl, msg, attrs...)
}
