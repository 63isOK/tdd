package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"jim": 20,
			"tom": 10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("return jim's score", func(t *testing.T) {
		request := newScoreRequest("jim")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("return tom's score", func(t *testing.T) {
		request := newScoreRequest("tom")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("return 404 when on missing players", func(t *testing.T) {
		request := newScoreRequest("hugo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseCode(t, response.Code, http.StatusNotFound)
	})

	t.Run("accepted on post", func(t *testing.T) {
		request := newStoreRequest("cobra")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseCode(t, response.Code, http.StatusAccepted)
	})
}

func newScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newStoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func assertResponseCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
