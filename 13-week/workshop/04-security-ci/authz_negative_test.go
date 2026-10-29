package httpapi

// Negative authorization test for the trust boundary api-gateway -> clinical-service.
// Story: code-corhuila/dlc-docs#137 (STRIDE E, risk 9, P0).
// Copy into dlc-clinical-api/internal/adapter/in/httpapi/ and run:
//   go test ./internal/adapter/in/httpapi/ -run TestClinicalRecordAuthz -v

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"testing"

	"github.com/code-corhuila/dlc-clinical-api/internal/adapter/out/persistence"
	"github.com/code-corhuila/dlc-clinical-api/internal/application/usecase"
)

const unassignedDentist = "33333333-3333-4333-8333-333333333333"

// Criterion 1: a DENTIST token that is not assigned to the patient cannot read
// GET /api/v1/clinical-records/{id}. The response must not be 200 and must not
// reveal the record (404 preferred, 403 accepted).
func TestClinicalRecordAuthzUnassignedDentistCannotRead(t *testing.T) {
	h := recordsRouter(stubCare{}, stubPatients{})
	created := post(h, validToken(t), key1, `{"patientId":"`+patientA+`"}`)
	if created.Code != http.StatusCreated && created.Code != http.StatusOK {
		t.Fatalf("setup: create record = %d %s", created.Code, created.Body.String())
	}
	var rec map[string]any
	_ = json.Unmarshal(created.Body.Bytes(), &rec)
	id, _ := rec["id"].(string)

	got := do(h, http.MethodGet, "/api/v1/clinical-records/"+id, tokenFor(t, unassignedDentist, "DENTIST"), nil, nil)
	if got.Code != http.StatusNotFound && got.Code != http.StatusForbidden {
		t.Fatalf("unassigned dentist read = %d, want 404 or 403 (%s)", got.Code, got.Body.String())
	}
	if got.Code == http.StatusForbidden {
		t.Logf("returns 403: record existence is revealed, see criterion 1 of dlc-docs#137 (want 404)")
	}
}

// Criterion 3: a SECRETARY_ASSISTANT token cannot open a clinical record.
func TestClinicalRecordAuthzAssistantCannotOpenRecord(t *testing.T) {
	h := recordsRouter(stubCare{}, stubPatients{})
	got := post(h, tokenFor(t, assistant, "SECRETARY_ASSISTANT"), key1, `{"patientId":"`+patientA+`"}`)
	if got.Code != http.StatusForbidden {
		t.Fatalf("assistant POST /api/v1/clinical-records = %d, want 403 (%s)", got.Code, got.Body.String())
	}
}

// switchCare answers like stubCare until down is set; then the Appointments
// assignment check fails, as if the service were unavailable.
type switchCare struct{ down *bool }

func (s switchCare) IsAssigned(ctx context.Context, p, d string) (bool, error) {
	if *s.down {
		return false, errors.New("appointments unavailable")
	}
	return stubCare{}.IsAssigned(ctx, p, d)
}
func (s switchCare) AssignedPatients(ctx context.Context, d string) ([]string, error) {
	if *s.down {
		return nil, errors.New("appointments unavailable")
	}
	return stubCare{}.AssignedPatients(ctx, d)
}

// Criterion 5 (review by @juan0898 on dlc-docs#133): if the assignment check
// fails, even the assigned dentist is denied (fail closed).
func TestClinicalRecordAuthzFailsClosedWhenAppointmentsIsDown(t *testing.T) {
	down := false
	care := switchCare{down: &down}
	store := persistence.NewMemory()
	svc := &usecase.RecordService{Repo: store, Care: care, Patients: stubPatients{}, Guard: testGuard(store, care, stubPatients{}), Clock: &clock{t: testNow}, IDs: uuids{}}
	h := NewRouter(Options{Verifier: testVerifier(), Logger: slog.New(slog.NewJSONHandler(io.Discard, nil)), Routes: RecordRoutes(svc)})

	created := post(h, validToken(t), key1, `{"patientId":"`+patientA+`"}`)
	var rec map[string]any
	_ = json.Unmarshal(created.Body.Bytes(), &rec)
	id, _ := rec["id"].(string)

	down = true
	got := do(h, http.MethodGet, "/api/v1/clinical-records/"+id, validToken(t), nil, nil)
	if got.Code == http.StatusOK {
		t.Fatalf("assigned dentist read with Appointments down = 200, want denied (%s)", got.Body.String())
	}
	t.Logf("Appointments down -> %d", got.Code)
}
