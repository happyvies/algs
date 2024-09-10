package kata

// https://www.codewars.com/kata/58ca658cc0d6401f2700045f/train/go

func FindMultiples(integer, limit int) []int {
	count := integer
	args := []int{}

	for count <= limit {
		if count%integer == 0 {
			args = append(args, count)
		}

		count++
	}

	return args
}
