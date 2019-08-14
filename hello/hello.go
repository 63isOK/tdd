package main

import "fmt"

const helloPrefix = "hello "

// HelloWorld say hello to me
func HelloWorld(who string) string {
	return helloPrefix + who
}

func main() {
	fmt.Println(HelloWorld("Tom"))
}
