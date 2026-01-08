package main

import (
	"fmt"
	"regexp"
)

func ValidateEmail() {
	var email string
	fmt.Scanln(&email)

	validate(email)
}

func validate(email string) {
	pattern := `^[a-zA-z0-9._+-]+@[a-zA-z0-9]+\.[a-zA-z]{2,5}$`
	emailPattern := regexp.MustCompile(pattern)

	if emailPattern.MatchString(email) {
		fmt.Println("Valid email address")
	} else {
		fmt.Println("Invalid email address")
	}
}
