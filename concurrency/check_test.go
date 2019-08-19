package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://www.abc.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.baidu.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.abc.com",
	}

	result := CheckWebsites(mockWebsiteChecker, websites)
	want := len(websites)
	got := len(result)

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"https://www.baidu.com":    true,
		"https://www.google.com":   true,
		"https://www.facebook.com": true,
		"https://www.abc.com":      false,
	}

	if !reflect.DeepEqual(expectedResults, result) {
		t.Fatalf("want %v, got %v", expectedResults, result)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "hello"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
