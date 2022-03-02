package main

import "fmt"

const englishHelloPrefix = "Hello"

const spanishHelloPrefix = "Hola"

const frenchHelloPrefix = "Bonjour"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := gettingPrefix(language)

	return fmt.Sprintf("%s, %s", prefix, name)
}

func gettingPrefix(language string) string {
	switch language {
	case "Spanish":
		return spanishHelloPrefix
	case "French":
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("Khanh", "Spanish"))
}
