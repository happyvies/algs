package main

import (
	"sort"
)

// https://www.codewars.com/kata/5656b6906de340bd1b0000ac

func main() {
	TwoToOne("aretheyhere", "yestheyarehere")
	TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding")
}

func TwoToOne(s1 string, s2 string) string {
	fullString := s1 + s2
	uniqCharcs := make(map[rune]struct{}, 0)
	ans := make([]rune, 0)

	// Searching for unique values
	for _, charc := range fullString {
		if _, exist := uniqCharcs[charc]; !exist {
			uniqCharcs[charc] = struct{}{}
			ans = append(ans, charc)
		}
	}

	// Sorting values in slice
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})

	return string(ans)
}
