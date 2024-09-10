package kata

import (
	"strings"
)

// https://www.codewars.com/kata/57a1fd2ce298a731b20006a4/train/go

func IsPalindrome(str string) bool {
	args := []byte(strings.ToLower(str))

	for i := 0; i < len(args)/2; i++ {
		if string(args[i]) != string(args[len(args)-1-i]) {
			return false
		}
	}
	return true
}
