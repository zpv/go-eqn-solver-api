package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNoEquation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	body := bytes.NewBuffer([]byte("{\"eqn\": \"\"}"))

	req, err := http.NewRequest("POST", "/solve", body)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("Post request generation failed %d", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 400 {
		t.Error("Expected 400 error.")
	}
}
