package main

import (
	"fmt"
	"regexp"
)

func ValidateNumber() {
	// numbers := []string{"(123) 456-7890", "123-456-7890", "123.456.7890", "123-4567"}
	numbers := []string{"1234567890", "123456789"}
	for _, number := range numbers {
		matchTenNums(number)
	}

	separators := []string{"1234567890", "123-456-7890", "123.456.7890", "123 456 7890", "123-456-789"}
	for _, number := range separators {
		matchSep(number)
	}

	paretheses := []string{"(123) 456-7890", "(123)456-7890", "(123)-456-7890", "123.456.7890", "((123) 456-7890"}
	for _, number := range paretheses {
		matchPar(number)
	}
}

func checkRegexNumber(regex *regexp.Regexp, number string) {
	if regex.MatchString(number) {
		fmt.Println(number, "is a phone number")
	} else {
		fmt.Println(number, "is not a phone number")
	}
}

// match 10 digits
func matchTenNums(number string) {
	pattern := `[0-9]{10}`
	numPattern := regexp.MustCompile(pattern)

	checkRegexNumber(numPattern, number)
}

// 10 digits + separators
func matchSep(number string) {
	pattern := `\d{3}[- .]?\d{3}[- .]?\d{4}`
	numPattern := regexp.MustCompile(pattern)

	checkRegexNumber(numPattern, number)
}

// paretheses around area code
func matchPar(number string) {
	pattern := regexp.MustCompile(`^(\(\d{3}\)|\d{3})[- .]?\d{3}[- .]?\d{4}$`)

	checkRegexNumber(pattern, number)
}
