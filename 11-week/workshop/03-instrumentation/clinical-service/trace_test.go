package clinicalobs

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

const incoming = "00-7a1c2e9f04b84d3e9c1f6a0b2d4e8f13-b3f1a07c9d2e4a51-01"

// completionHandler stands in for POST /api/v1/procedures/{id}/completions:
// it records the outbox headers and calls Billing over HTTP with the same trace.
func completionHandler(log *slog.Logger, m *Metrics, outbox map[string]string, billingURL string) http.Handler {
	client := &http.Client{Transport: TracingTransport{}}
	mux := http.NewServeMux()
	mux.Handle("GET /metrics", m)
	mux.HandleFunc("POST /api/v1/procedures/{id}/completions", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		InjectMessageHeaders(r.Context(), outbox)
		req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, billingURL+"/api/v1/procedure-prices/OPER-01", nil)
		resp, err := client.Do(req)
		if err != nil {
			m.Observe("error", time.Since(start))
			LogEvent(r.Context(), log, slog.LevelWarn, "billing.price.timeout", "billing-service did not answer", time.Since(start))
			http.Error(w, "unavailable", http.StatusServiceUnavailable)
			return
		}
		resp.Body.Close()
		m.Observe("accepted", time.Since(start))
		LogEvent(r.Context(), log, slog.LevelInfo, "procedure.completion.accepted", "procedure completion committed with outbox event",
			time.Since(start), "procedure_id", r.PathValue("id"), "procedure_code", "OPER-01", "closure_status", "CLOSURE_PENDING")
		w.WriteHeader(http.StatusAccepted)
	})
	return WithTrace(mux)
}

func TestTraceIDReceivedEqualsTraceIDSentDownstream(t *testing.T) {
	var billingHeader string
	billing := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		billingHeader = r.Header.Get("traceparent")
	}))
	defer billing.Close()

	var logs bytes.Buffer
	log := NewLogger(&logs, "clinical-service", "2.0.0", "dev")
	m := NewMetrics()
	outbox := map[string]string{}
	srv := httptest.NewServer(completionHandler(log, m, outbox, billing.URL))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodPost, srv.URL+"/api/v1/procedures/5d0e8a7b-3c21-4f6e-9a10-2b7c4d8e1f05/completions", nil)
	req.Header.Set("traceparent", incoming)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	want, _, _ := ParseTraceparent(incoming)
	for name, h := range map[string]string{"billing http": billingHeader, "outbox": outbox["traceparent"], "response": resp.Header.Get("traceparent")} {
		got, _, ok := ParseTraceparent(h)
		if !ok || got != want {
			t.Errorf("%s: trace_id %q, want %q", name, got, want)
		}
	}

	var line map[string]any
	if err := json.Unmarshal(logs.Bytes(), &line); err != nil {
		t.Fatalf("log is not one JSON line: %v", err)
	}
	for _, k := range []string{"timestamp", "level", "service", "version", "environment", "trace_id", "span_id", "correlation_id", "event", "message", "duration_ms"} {
		if _, ok := line[k]; !ok {
			t.Errorf("log line misses %q", k)
		}
	}
	if line["trace_id"] != want {
		t.Errorf("log trace_id %v, want %s", line["trace_id"], want)
	}
	os.Stdout.Write(logs.Bytes())

	mr, _ := http.Get(srv.URL + "/metrics")
	var body bytes.Buffer
	body.ReadFrom(mr.Body)
	mr.Body.Close()
	if !strings.Contains(body.String(), `clinical_procedure_completions_total{outcome="accepted"} 1`) {
		t.Errorf("metrics missing accepted counter:\n%s", body.String())
	}
}

func TestInvalidTraceparentStartsNewTrace(t *testing.T) {
	if _, _, ok := ParseTraceparent("00-00000000000000000000000000000000-b3f1a07c9d2e4a51-01"); ok {
		t.Fatal("all-zero trace id must be rejected")
	}
}
