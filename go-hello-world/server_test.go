package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	e := echo.New()
	want := "Hello, World!"
	req := httptest.NewRequest(http.MethodGet, "/hello-world", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, sayHello(c)) {
		assert.Equalf(t, http.StatusOK, rec.Code, "expecting status  to be %d but got %d", http.StatusOK, rec.Code)
		got := strings.TrimSpace(rec.Body.String())
		assert.Equalf(t, want, got, "want %s but got %s", want, got)
	}
}
