package main

import (
	"flag"
	"fmt"
)

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

func greet(lang language) string {
	greeting, ok := phrasebook[lang]
	if !ok {
		return fmt.Sprintf("Unsupported language: %q", lang)
	}
	return greeting
}
