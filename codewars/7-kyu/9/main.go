package main

import (
	"fmt"
	"strings"
)

// https://www.codewars.com/kata/52fba66badcd10859f00097e

func main() {
	consonantLetters := Disemvowel("This website is for losers LOL!")
	fmt.Println(consonantLetters)
}

func Disemvowel(comment string) string {
	vowelLetters := map[string]struct{}{
		"a": {},
		"A": {},
		"e": {},
		"E": {},
		"i": {},
		"I": {},
		"o": {},
		"O": {},
		"u": {},
		"U": {},
	}
	chars := strings.Split(comment, "")

	consonantLetters := make([]string, 0)
	for _, char := range chars {
		if _, exist := vowelLetters[char]; !exist {
			consonantLetters = append(consonantLetters, char)
		}
	}

	return strings.Join(consonantLetters, "")
}
