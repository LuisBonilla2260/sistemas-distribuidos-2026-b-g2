package clinicalobs

import (
	"context"
	"net/http"
)

// InjectMessageHeaders adds traceparent to the ProcedureCompleted outbox document
// headers, so the relay can copy them into the AMQP message after the request ended.
func InjectMessageHeaders(ctx context.Context, headers map[string]string) {
	tc, ok := FromContext(ctx)
	if !ok {
		return
	}
	headers["traceparent"] = tc.NewChild().Traceparent()
	if id, ok := ctx.Value(corrKey{}).(string); ok {
		headers["x-correlation-id"] = id
	}
}

// TracingTransport propagates traceparent on outgoing HTTP calls.
type TracingTransport struct{ Base http.RoundTripper }

func (t TracingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	tc, ok := FromContext(r.Context())
	if !ok {
		return t.base().RoundTrip(r)
	}
	r = r.Clone(r.Context())
	r.Header.Set("traceparent", tc.NewChild().Traceparent())
	return t.base().RoundTrip(r)
}

func (t TracingTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}
