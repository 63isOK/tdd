package website

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createHTTPServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {

	slowServer := createHTTPServer(20 * time.Millisecond)
	fastServer := createHTTPServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()

	url1 := slowServer.URL
	url2 := fastServer.URL

	want := url2
	got := racer(url1, url2)

	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

}
