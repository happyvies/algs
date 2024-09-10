package kata

// https://www.codewars.com/kata/55ecd718f46fba02e5000029/train/go

func Between(a, b int) []int {
	args := []int{}

	for a <= b {
		args = append(args, a)
		a++
	}

	return args
}
