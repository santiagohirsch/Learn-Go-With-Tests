package main

import "fmt"

const enPrefix = "Hello "
const esPrefix = "Hola "
const frPrefix = "Bonjour "
const suffix = "!"
const spanish = "Spanish"
const french = "French"

func HelloWorldPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = esPrefix
	case french:
		prefix = frPrefix
	default:
		prefix = enPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return HelloWorldPrefix(language) + name + suffix
}

func main() {
	fmt.Println(Hello("World", "English"))
}
