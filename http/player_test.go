package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("return score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/socre", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		got := response.Body.String()
		want := "20"

		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}
