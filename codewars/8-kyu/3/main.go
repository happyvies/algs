package kata

// https://www.codewars.com/kata/5ce9c1000bab0b001134f5af/train/go

func QuarterOf(month int) int {
	switch {
	case month <= 3:
		return 1
	case month > 3 && month <= 6:
		return 2
	case month > 6 && month <= 9:
		return 3
	case month > 9 && month <= 12:
		return 4
	default:
		return 0
	}
}
