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

	prefix := englishPrefix

	switch language {
	case chineseLanguage:
		prefix = chinesePrefix
	case frenchLanguage:
		prefix = frenchPrefix
	}

	return prefix + who
}

func main() {
	fmt.Println(HelloWorld("Tom", ""))
}
