package main

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNoEquation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	body := bytes.NewBuffer([]byte("{\"eqn\": \"\"}"))

	req, err := http.NewRequest("POST", "/solve", nil)
	if err == nil {
		t.Errorf("Should have expected an error.", err)
	}
}
