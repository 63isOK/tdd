package main

import (
	"fmt"
)

func main() {

	fmt.Println("123", len("123"))
	fmt.Println("123沃尔夫", len("123沃尔夫"))
	fmt.Println("123abc", len("123abc"))
	fmt.Println("阿斯蒂芬 阿斯蒂芬", len("阿斯蒂芬 阿斯蒂芬"))

	fmt.Println("沃", len("沃"))
	fmt.Println("尔", len("尔"))
	fmt.Println("夫", len("夫"))

	s := "沃尔夫"
	fmt.Println(s, len(s))
	fmt.Println([]rune(s), len([]rune(s)))
}
