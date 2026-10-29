package httpapi

// Negative authorization test for the trust boundary api-gateway -> clinical-service.
// Story: code-corhuila/dlc-docs#137 (STRIDE E, risk 9, P0).
// Copy into dlc-clinical-api/internal/adapter/in/httpapi/ and run:
//   go test ./internal/adapter/in/httpapi/ -run TestClinicalRecordAuthz -v

import (
	"encoding/json"
	"net/http"
	"testing"
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
