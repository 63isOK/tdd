package loop

import (
	"strings"
)

func repeat(char string, count int) (ret string) {
	for i := 0; i < count; i++ {
		ret += char
	}

	return
}

// RepeatWithStd repeat string with std lib
func RepeatWithStd(char string, count int) string {
	if count < 0 {
		return ""
	}

	return strings.Repeat(char, count)
}
