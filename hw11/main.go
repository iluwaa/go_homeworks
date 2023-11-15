package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	path := "numbers.txt"
	findPhones(path)
	fmt.Println()
	path = "text.txt"

	searchPatterns(path)

}

func searchPatterns(path string) {
	text, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Original text:")
	fmt.Println(string(text))

	r := regexp.MustCompile(`\p{Lu}\p{Cyrillic}+\p{Ll}`)

	fmt.Println()
	fmt.Println("Words which starts from uppercase:")

	for _, word := range r.FindAllString(string(text), -1) {
		fmt.Println(word)
	}

	r = regexp.MustCompile(`(?:\s|$)([^\nАЕЄИІЇОУЮЯаеєиіїоуюя]\p{Cyrillic}+[АЕЄИІЇОУЮЯаеєиіїоуюя])(?:\s|,|.|$)`)

	fmt.Println()
	fmt.Println("Words which starts from vowel and ends with consonant:")

	for _, word := range r.FindAllStringSubmatch(string(text), -1) {
		fmt.Println(word[1])
	}
}

func findPhones(path string) {
	text, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Original text:")
	fmt.Println(string(text))
	fmt.Println()

	r := regexp.MustCompile(`\(?\d{3}\)?[\d\s-]?\d{3}[\d\s-]?\d{4}`)

	fmt.Println("Matched phone numbers:")

	for _, phone := range r.FindAllString(string(text), -1) {
		fmt.Println(phone)
	}
}
