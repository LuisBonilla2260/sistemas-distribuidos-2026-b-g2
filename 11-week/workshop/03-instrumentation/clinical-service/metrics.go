package clinicalobs

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

// Buckets in seconds: the Dentist waits with the patient in the chair, so 300 ms is the target.
var latencyBuckets = []float64{0.05, 0.1, 0.2, 0.3, 0.5, 1, 2.5}

// Metrics exposes clinical_procedure_completions_total{outcome} and
// clinical_procedure_completion_duration_seconds in Prometheus text format.
type Metrics struct {
	mu      sync.Mutex
	byOut   map[string]int
	buckets []int
	sum     float64
	count   int
}

func NewMetrics() *Metrics {
	return &Metrics{byOut: map[string]int{}, buckets: make([]int, len(latencyBuckets))}
}

func (m *Metrics) Observe(outcome string, d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.byOut[outcome]++
	s := d.Seconds()
	for i, b := range latencyBuckets {
		if s <= b {
			m.buckets[i]++
		}
	}
	m.sum += s
	m.count++
}

func (m *Metrics) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	m.mu.Lock()
	defer m.mu.Unlock()
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	fmt.Fprintln(w, "# TYPE clinical_procedure_completions_total counter")
	outs := make([]string, 0, len(m.byOut))
	for o := range m.byOut {
		outs = append(outs, o)
	}
	sort.Strings(outs)
	for _, o := range outs {
		fmt.Fprintf(w, "clinical_procedure_completions_total{outcome=%q} %d\n", o, m.byOut[o])
	}
	fmt.Fprintln(w, "# TYPE clinical_procedure_completion_duration_seconds histogram")
	for i, b := range latencyBuckets {
		fmt.Fprintf(w, "clinical_procedure_completion_duration_seconds_bucket{le=\"%g\"} %d\n", b, m.buckets[i])
	}
	fmt.Fprintf(w, "clinical_procedure_completion_duration_seconds_bucket{le=\"+Inf\"} %d\n", m.count)
	fmt.Fprintf(w, "clinical_procedure_completion_duration_seconds_sum %g\n", m.sum)
	fmt.Fprintf(w, "clinical_procedure_completion_duration_seconds_count %d\n", m.count)
}
