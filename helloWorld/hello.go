package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const yoruba = "Yoruba"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const yorubaHelloPrefix = "Sho da, "

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
	case yoruba:
		prefix = yorubaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	
	return
}

func main() {

	fmt.Println(Hello("world", ""))

	fmt.Println(Hello("world", french))

}
