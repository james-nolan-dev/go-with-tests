package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	irish              = "Irish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	irishHelloPrefix   = "Dia dhuit, "
)

func main() {
	fmt.Println(Hello("world", ""))
}

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
	case irish:
		prefix = irishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
