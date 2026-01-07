package main

import (
	"fmt"
	"regexp"
	"strings"
)

func extractURL() {
	input := "Visit us at https://example.com or http://test.com for more info. credits http://www.chat.py"
	urls := extractURLs(input)
	urlsRegex := ExtractURLsRegex(input)
	for _, url := range urls {
		println(url) // https://example.com, http://test.com
	}
	fmt.Println("using regex")
	for _, url := range urlsRegex {
		println(url) // https://example.com, www.chat.py
	}
}

// without regex
func extractURLs(input string) []string {
	var urls []string
	words := strings.Fields(input)

	for _, word := range words {
		if strings.HasPrefix(word, "http://") || strings.HasPrefix(word, "https://") || strings.HasPrefix(word, "www.") {
			urls = append(urls, word)
		}
	}

	return urls
}

// elegant less space but less readabale
func ExtractURLsRegex(input string) []string {
	pattern := regexp.MustCompile(`https://\S+|www\.\S+`)
	return pattern.FindAllString(input, -1)
}

