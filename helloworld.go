package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε!",     // Greek
	"en": "Hello, World!",      // English
	"es": "Hola, Mundo!",       // Spanish
	"fr": "Bonjour, le monde!", // French
	"he": "שלום עולם!",         // Hebrew
	"it": "Ciao, mondo!",       // Italian
	"ja": "こんにちは世界!",           // Japanese
	"ur": "ہیلو دنیا",          // Urdu
	"vi": "Xin chào Thế Giới",  // Vietnamese
	"zh": "你好，世界！",             // Chinese
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
