package concurrency

// WebsiteChecker is check tool for website
type WebsiteChecker func(string) bool

// CheckWebsites check urls
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
