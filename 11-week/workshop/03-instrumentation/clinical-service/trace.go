package clinicalobs

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"
)

// TraceContext is the W3C trace context of the current hop.
type TraceContext struct {
	TraceID  string
	SpanID   string
	ParentID string
}

type ctxKey struct{}

func FromContext(ctx context.Context) (TraceContext, bool) {
	tc, ok := ctx.Value(ctxKey{}).(TraceContext)
	return tc, ok
}

func randomHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// ParseTraceparent accepts "00-<32 hex>-<16 hex>-<2 hex>" and rejects all-zero ids.
func ParseTraceparent(h string) (traceID, parentID string, ok bool) {
	p := strings.Split(strings.TrimSpace(h), "-")
	if len(p) != 4 || p[0] != "00" || !isHex(p[1], 32) || !isHex(p[2], 16) || !isHex(p[3], 2) {
		return "", "", false
	}
	if strings.Trim(p[1], "0") == "" || strings.Trim(p[2], "0") == "" {
		return "", "", false
	}
	return p[1], p[2], true
}

func isHex(s string, n int) bool {
	if len(s) != n {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil && strings.ToLower(s) == s
}

func (tc TraceContext) Traceparent() string {
	return "00-" + tc.TraceID + "-" + tc.SpanID + "-01"
}

// NewChild starts a span in the same trace, used for outgoing calls and the outbox relay.
func (tc TraceContext) NewChild() TraceContext {
	return TraceContext{TraceID: tc.TraceID, SpanID: randomHex(8), ParentID: tc.SpanID}
}

// WithTrace reads traceparent or starts a new trace, stores it in the request
// context and echoes it plus X-Correlation-Id in the response.
func WithTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc := TraceContext{SpanID: randomHex(8)}
		if traceID, parent, ok := ParseTraceparent(r.Header.Get("traceparent")); ok {
			tc.TraceID, tc.ParentID = traceID, parent
		} else {
			tc.TraceID = randomHex(16)
		}
		w.Header().Set("traceparent", tc.Traceparent())
		corr := correlationID(r, tc)
		w.Header().Set("X-Correlation-Id", corr)
		ctx := context.WithValue(context.WithValue(r.Context(), ctxKey{}, tc), corrKey{}, corr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// correlationID keeps a safe incoming X-Correlation-Id, as dlc-clinical-api does, else uses the trace id.
func correlationID(r *http.Request, tc TraceContext) string {
	id := r.Header.Get("X-Correlation-Id")
	if id == "" || len(id) > 128 {
		return tc.TraceID
	}
	for _, c := range id {
		if c < 0x21 || c > 0x7e {
			return tc.TraceID
		}
	}
	return id
}
