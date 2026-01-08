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

	// Match any lowercase letter from 'a' to 'z'
	pattern := regexp.MustCompile("[a-z]")
	lowercaseLetters := pattern.FindAllString(text, -1)
	fmt.Println("Lowercase letters:", lowercaseLetters)

	// Match any word character (alphanumeric + underscore)
	pattern = regexp.MustCompile("[A-Za-z0-9_]")
	wordCharacters := pattern.FindAllString(text, -1)
	fmt.Println("Word characters:", wordCharacters)

	// Match any character that is NOT a space
	pattern = regexp.MustCompile("[^ ]")
	nonSpaceCharacters := pattern.FindAllString(text, -1)
	fmt.Println("Non-space characters:", nonSpaceCharacters)
}

func AsterixExample() {
	// one or more 'o' then k
	pattern := regexp.MustCompile(`o*k`)
	fmt.Println(pattern.FindString("oogenesis")) // returns ""
	fmt.Println(pattern.FindString("cook"))      // returns "ook"
	fmt.Println(pattern.FindString("oocyst"))    // returns ""
	fmt.Println(pattern.FindString("book"))      // returns "ook"
}

func PlusExample() {
	// one g and one+ o
	pattern := regexp.MustCompile(`go+`)
	fmt.Println(pattern.FindAllString("goo", -1))   // returns [goo]
	fmt.Println(pattern.FindAllString("goooo", -1)) // returns [goooo]
	fmt.Println(pattern.FindAllString("g", -1))     // returns []
	fmt.Println(pattern.FindAllString("gogoo", -1)) // returns [go goo]
}

func CurlyExample() {
	// 2 - 4 o
	pattern := regexp.MustCompile(`go{2,4}`)
	fmt.Println(pattern.FindAllString("go", -1))        // returns []
	fmt.Println(pattern.FindAllString("goo", -1))       // returns [goo]
	fmt.Println(pattern.FindAllString("goooooooo", -1)) // returns [goooo]
}

func LazyQuantifiers() {
	pattern := regexp.MustCompile(`".*?"`)               // search for minimal match
	fmt.Println(pattern.FindString(`"foo" "bar" "baz"`)) // returns "foo"
}

func Repetition() {
	// parentheses
	text := `The detective carefully examined the crime scene. It was a gruesome sight—the victim had been brutally killed. The evidence pointed towards a professional hitman as the perpetrator. The investigators were determined to solve the case and bring the killer to justice. The motive behind the killing remained unclear, but the police were determined to uncover the truth.`
	match := regexp.MustCompile(`(kill)+`) // why plus if replace all? bad use case
	modStr := match.ReplaceAll([]byte(text), []byte("****"))
	fmt.Println(string(modStr))

	// dot
	pattern := regexp.MustCompile(`h.t`)
	fmt.Println(pattern.MatchString("hat")) // true
	fmt.Println(pattern.MatchString("hot")) // true
	fmt.Println(pattern.MatchString("bit")) // false
	pattern = regexp.MustCompile("h.*t")
	fmt.Println(pattern.FindString("have you got money and bitches")) // have you got money and bit

	// pipe
	pattern = regexp.MustCompile(`dog|cat`)
	fmt.Println(pattern.MatchString("dogs")) // true
	fmt.Println(pattern.MatchString("cat"))  // true
	fmt.Println(pattern.MatchString("bird")) // false

}

func Escape() {
	re := regexp.MustCompile(`\d+\.\d+`)
	fmt.Println(re.MatchString("99.9"))  // true
	fmt.Println(re.MatchString("9999.")) // false
}

func Anchors() {
	re := regexp.MustCompile("^start.*end$")
	fmt.Println(re.MatchString("start at the end"))    // true
	fmt.Println(re.MatchString("start at the bottom")) // false
}

func WordBoundaries() {
	// start
	pattern := regexp.MustCompile(`(?i)\bbet`)
	fmt.Println(pattern.FindAllString("Betty's better bet was to buy the blue blouse.", -1)) // [Bet bet, bet]

	// end
	re := regexp.MustCompile(`(?i)sion\b`)
	fmt.Println(re.FindAllString("After much discussion, the team came to a consensus on the vision for the project.", -1)) // [sion sion]

	// entire word
	re = regexp.MustCompile(`(?i)\bcat\b`)
	fmt.Println(re.FindAllString("Cat and dog and dogs and cats", -1)) // [Cat]

}

func Grouping() {
	pattern := regexp.MustCompile(`(quick|lazy) (brown|gray) (fox|dog)`)
	text := "The lazy gray dog jumps over the lazy fox"
	fmt.Println(pattern.FindAllString(text, -1)) // returns [lazy gray dog]

	// rearrange day and month
	pattern = regexp.MustCompile(`(\d{2})-(\d{2})-(\d{4})`)
	fmt.Println(pattern.ReplaceAllString("12-31-2025", "$2-$1-$3"))
}

func MatchingUnicode() {
	text := "Hello world! 世界你好！"

	re := regexp.MustCompile(`\p{Han}`)
	fmt.Println(re.FindAllString(text, -1))
}
