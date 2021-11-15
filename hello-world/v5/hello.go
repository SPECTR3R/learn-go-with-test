package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	french             = "French"
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
)

func Hello(name, lang string) string {
	prefix := greetingPrefix(lang)
	if name == "" {
		return prefix + "World"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Juan", ""))
}
