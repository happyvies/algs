package kata

import "strings"

// 8 kyu
// https://www.codewars.com/kata/57e76bc428d6fbc2d500036d/train/go

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}
