package website

import (
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}

func racer(urlA, urlB string) string {

	costA := measureResponseTime(urlA)
	costB := measureResponseTime(urlB)

	if costA < costB {
		return urlA
	}

	return urlB
}
