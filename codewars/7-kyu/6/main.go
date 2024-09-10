package kata

// https://www.codewars.com/kata/563b662a59afc2b5120000c6/train/go

func NbYear(start int, percent float64, aug int, final int) int {
	var years int
	var increase float64
	percent = percent / 100

	for start < final {
		increase = float64(start) * percent
		start = start + int(increase) + aug

		years++
	}

	return years
}
