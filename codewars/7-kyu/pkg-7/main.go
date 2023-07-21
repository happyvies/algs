package kata

// 7 kyu
// https://www.codewars.com/kata/62c93765cef6f10030dfa92b/train/go

func Cats(start, finish int) int {
	if start == finish {
		return 0
	}

	var steps int

	for start < finish {
		if finish-start < 3 {
			steps++
			start++
			continue

		}

		steps++
		start += 3
	}

	return steps
}
