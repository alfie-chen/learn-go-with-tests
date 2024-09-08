package main

import "fmt"

const (
	chinese = "Chinese"
	french = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	chineseHelloPrefix = "你好, "
)


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case french:
			prefix = frenchHelloPrefix
		case spanish:
			prefix = spanishHelloPrefix
		case chinese:
			prefix = chineseHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Alfie", "Spanish"))
}