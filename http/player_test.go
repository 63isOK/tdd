package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("return jim's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/jim", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		got := response.Body.String()
		want := "20"

		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("return tom's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/tom", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		got := response.Body.String()
		want := "10"

		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}
