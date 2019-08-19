package website

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	url1 := slowServer.URL
	url2 := fastServer.URL

	want := url2
	got := racer(url1, url2)

	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	slowServer.Close()
	fastServer.Close()
}
