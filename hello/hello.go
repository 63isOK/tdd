package main

import "fmt"

const (
	englishPrefix   = "hello "
	englishLanguage = "english"
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

	return getPrefix(language) + who
}

func getPrefix(language string) (prefix string) {
	switch language {
	case chineseLanguage:
		prefix = chinesePrefix
	case frenchLanguage:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func main() {
	fmt.Println(HelloWorld("Tom", ""))
}
