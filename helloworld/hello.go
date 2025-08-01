// Package helloworld provides functions for greeting people in different languages
package helloworld

const (
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	spanishLanguage     = "Spanish"
	frenchHelloPrefix   = "Bonjour, "
	frenchLanguage      = "French"
	japaneseLanguage    = "Japanese"
	japaneseHelloPrefix = "こんにちは, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	case japaneseLanguage:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
