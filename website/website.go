package website

import (
	"net/http"
)

func racer(urlA, urlB string) string {
	select {
	case <-ping(urlA):
		return urlA
	case <-ping(urlB):
		return urlB
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
