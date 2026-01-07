package main

import (
	"fmt"
	"regexp"
)

// litaral patterns
// match precisely match characters

func literalPatternExample() {
	pattern := regexp.MustCompile("Kain")
	text := "My Name is Artem Kain currently i work on a pet project"

	if pattern.MatchString(text) {
		fmt.Println("Text is about me")
	} else {
		fmt.Println("Text is not about me")
	}
}

// flags are used for more complex patterns
func flagsExample() {
	// i flag (ignore)
	text := "Apples are delicious, but so are apples, aPpleS, and APPlE."
	pattern := regexp.MustCompile(`(?i)apple`)
	matches := pattern.FindAllString(text, -1)
	fmt.Println(matches) // [Apple apple aPple APPlE]

	// g flag (global) (in go implicated with findAllStrings method)

	// m flag (multiline)
	text = "Hello World\nHello Golang\nHello Lord Kain"
	pattern = regexp.MustCompile(`(?m)^Hello`)
	matches = pattern.FindAllString(text, -1)
	fmt.Println(matches)

	// s flag (single-line) useful if pattern is located in multiple lines
	text = "Pattern\nspanning\nmultiple lines.\nExample."
	pattern = regexp.MustCompile(`(?s)spanning.*Example`) //[spanning; example]
	match := pattern.FindString(text)
	fmt.Println(match)
}

func charClassesExample() {
	// \w, \d
	// find all words containing at least one digit
	text := "order will be 120.000$ and your confimation code is CD12TX"
	pattern := regexp.MustCompile(`\w+\d\w*`)
	matches := pattern.FindAllString(text, -1)
	fmt.Println(matches)
}

func NegatedCharSet() {
	text := "I will be back!"
	noVowelPattern := regexp.MustCompile(`[^aeiouAEIOU\s]`)
	matches := noVowelPattern.FindAllString(text, -1)
	fmt.Println(matches)
}

func RangesExample() {
	text := "The numbers are 5, 25, 49, and 50."
	tenToFourtyPattern := regexp.MustCompile(`[1-4][0-9]`)
	matches := tenToFourtyPattern.FindAllString(text, -1)
	fmt.Println(matches) // 25, 49
}
