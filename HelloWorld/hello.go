package main

import "fmt"

const helloWorldPrefix = "Hello "
const helloWorldSuffix = "!"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloWorldPrefix + name + helloWorldSuffix
}

func main() {
	fmt.Println(Hello("World"))
}
