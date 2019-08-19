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

	t.Run("without time", func(t *testing.T) {
		slowServer := createHTTPServer(20 * time.Millisecond)
		fastServer := createHTTPServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		url1 := slowServer.URL
		url2 := fastServer.URL

		want := url2
		got, err := racer(url1, url2)

		if err != nil {
			t.Fatalf("didn't excepted an error, but got one %v", err)
		}

		if want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	})

	t.Run("test with timeout", func(t *testing.T) {
		server := createHTTPServer(15 * time.Millisecond)
		defer server.Close()

		url := server.URL

		_, err := configurableRacer(url, url, 10*time.Millisecond)

		if err == nil {
			t.Errorf("expected a error, but didn't get one")
		}
	})
}
