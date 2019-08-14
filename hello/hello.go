package main

import "fmt"

const (
	defaultPrefix   = "hello "
	chinesePrefix   = "你好 "
	chineseLanguage = "chinese"
	frenchPrefix    = "bonjour "
	frenchLanguage  = "french"
)

// HelloWorld say hello to me
func HelloWorld(who, language string) string {
	if who == "" {
		who = "world"
	}

	if language == chineseLanguage {
		return chinesePrefix + who
	}

	if language == frenchLanguage {
		return frenchPrefix + who
	}

	return defaultPrefix + who

}

func main() {
	fmt.Println(HelloWorld("Tom", ""))
}
