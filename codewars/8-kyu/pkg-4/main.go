package kata

import (
	"strconv"
	"strings"
)

// 8 kyu
// https://www.codewars.com/kata/57eae65a4321032ce000002d/train/go

func FakeBin(x string) string {
	var sb strings.Builder

	for i := 0; i < len(x); i++ {
		num, _ := strconv.Atoi(string(x[i]))

		if num < 5 {
			sb.WriteByte('0')
		} else {
			sb.WriteByte('1')
		}
	}

	return sb.String()
}
