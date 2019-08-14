package main

import "fmt"

// HelloWorld say hello to me
func HelloWorld(who string) string {
	return fmt.Sprintf("hello %s", who)
}

func main() {
	fmt.Println(HelloWorld("Tom"))
}
