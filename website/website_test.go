package website

import "testing"

func TestRacer(t *testing.T) {
	url1 := "http://www.baidu.com"
	url2 := "http://www.bing.com"

	want := url1
	got := racer(url1, url2)

	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}
