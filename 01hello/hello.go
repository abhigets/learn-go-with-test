package main

import "fmt"

func main() {
	fmt.Println(Hello("Chris", "English"))
}

const spanish = "Spanish"
const french = "French"
const englishPrefixWithHello = "Hello, "
const spanishPrefixWithHola = "Hola, "
const frenchPrefixWithBonjour = "Bonjour, "


func Hello(name, language string) string {

	if name == "" {
		name = "World!"
	}

	prefix := ""
	switch language {
	case french:
		prefix = frenchPrefixWithBonjour
	case spanish:
		prefix = spanishPrefixWithHola
	default:
		prefix = englishPrefixWithHello

	}

	return prefix + name
}
