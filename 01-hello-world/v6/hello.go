package main

import "fmt"

const french = "fr"
const spanish = "es"
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const defaultHelloPrefix = "Hello, "

func main() {
	fmt.Println("Hello", "es")
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return GreetingPrefix(language) + name
}

func GreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		{
			prefix = frenchHelloPrefix
		}
	case spanish:
		{
			prefix = spanishHelloPrefix
		}
	default:
		prefix = defaultHelloPrefix
	}

	return prefix
}
