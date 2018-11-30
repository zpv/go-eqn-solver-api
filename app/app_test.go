package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestBasicEquation(t *testing.T) {
	testEquation(t, "2x=4", "x=2")
}

func TestAdvancedEquation(t *testing.T) {
	testEquation(t, "6x+4x+2x+1x=4x+5x+3", "x=0.75")
}

func TestNoSolutionEquation(t *testing.T) {
	testEquation(t, "x+5=x", "No solution")
}

func TestNoEquation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	resp := performSolveRequest(testRouter, "")

	if resp.Code != http.StatusBadRequest {
		t.Errorf("Unexpected response code: got %v want %v", resp.Code, http.StatusBadRequest)
	}
}

func TestBadlyFormattedEquation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	resp := performSolveRequest(testRouter, "=")

	if resp.Code != http.StatusBadRequest {
		t.Errorf("Unexpected response code: got %v want %v", resp.Code, http.StatusBadRequest)
	}
}

// testEquation is a test helper which validates if a given equation
// resolves to an expected string
func testEquation(t *testing.T, eqn, expected string) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	resp := performSolveRequest(testRouter, eqn)

	// Check status code is what we expect
	if resp.Code != http.StatusOK {
		t.Errorf("Unexpected response code: got %v want %v", resp.Code, http.StatusOK)
	}

	// Check response body is what we expect
	expectedMap := map[string]string{"result": expected}
	expectedJSON, _ := json.Marshal(expectedMap)

	expectedJSONString := string(expectedJSON)
	if resp.Body.String() != expectedJSONString {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Body.String(), expectedJSONString)
	}
}

// performSolveRequest takes in a linear equation as a string and
// returns a request to `/solve` with the equation as the body
func performSolveRequest(r http.Handler, eqn string) *httptest.ResponseRecorder {
	// Build JSON object
	eqnMap := map[string]string{"eqn": eqn}
	json, _ := json.Marshal(eqnMap)

	body := bytes.NewBuffer([]byte(string(json)))
	req, _ := http.NewRequest("POST", "/api/solve", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}
