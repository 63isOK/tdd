package website

import (
	"net/http"
	"time"
)

func racer(urlA, urlB string) string {
	startA := time.Now()
	_, _ = http.Get(urlA)
	costA := time.Since(startA)

	startB := time.Now()
	_, _ = http.Get(urlB)
	costB := time.Since(startB)

	if costA < costB {
		return urlA
	}

	return urlB
}
