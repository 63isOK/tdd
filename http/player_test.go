package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("return jim's score", func(t *testing.T) {
		request := newScoreRequest("jim")
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("return tom's score", func(t *testing.T) {
		request := newScoreRequest("tom")
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
