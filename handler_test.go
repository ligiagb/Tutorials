package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerURLNotFound(t *testing.T) {

	req := httptest.NewRequest("GET", "http://test.test/urlinfo/1/google.com/index.html", nil)
	w := httptest.NewRecorder()
	MalwareResponse(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

}

func TestHandlerURLFound(t *testing.T) {

	req := httptest.NewRequest("GET", "http://test.test/urlinfo/1/capacitacion.inami.gob.mx/capa.exe", nil)
	w := httptest.NewRecorder()
	MalwareResponse(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"bad"}
`, w.Body.String())

}
