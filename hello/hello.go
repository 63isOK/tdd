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

	switch language {
	case chineseLanguage:
		return chinesePrefix + who
	case frenchLanguage:
		return frenchPrefix + who
	case englishLanguage:
		fallthrough
	default:
		return englishPrefix + who
	}

}

func main() {
	fmt.Println(HelloWorld("Tom", ""))
}
