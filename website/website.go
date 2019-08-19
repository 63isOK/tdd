package website

import (
	"fmt"
	"net/http"
	"time"
)

var (
	timeout10s = 10 * time.Second
)

func racer(a, b string) (winner string, err error) {
	return configurableRacer(a, b, timeout10s)
}

func configurableRacer(urlA, urlB string, timeout time.Duration) (string, error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout waiting for %s and %s", urlA, urlB)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		_, _ = http.Get(url)
		ch <- true
	}()
	return ch
}
