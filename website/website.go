package website

import (
	"fmt"
	"net/http"
	"time"
)

func racer(urlA, urlB string) (string, error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(10 * time.Second):
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
