package kata

import (
	"strings"
)

// https://www.codewars.com/kata/59d9ff9f7905dfeed50000b0/train/go

const start int = 96

func solve(slice []string) []int {
	var positionAlphabet, positionCurrent, counter int = 0, 1, 0
	var element byte
	var str string = strings.ToLower(strings.Join(slice, " "))
	var args []int

	for i := 0; i < len(str); i++ {

		if str[i] == ' ' {
			args = append(args, counter)
			positionCurrent = 1
			counter = 0
			continue
		}

		element = str[i]
		positionAlphabet = int(element) - start

		if positionCurrent == positionAlphabet {
			counter++
		}
		positionCurrent++

		if i == len(str)-1 {
			args = append(args, counter)
		}

	}

	return args
}
