package main

import (
	"flag"
	"fmt"
)

func main() {
	// hold the value provided by the user
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur...")
	// to scan the input parameters
	flag.Parse()
	// test it with: go run main.go -lang=el

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
	"ca": "Hola món!",         // Catalan
}

// greet says hello to the world in the specified language
func greet(l language) string {
	// ok -> if exists...
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
