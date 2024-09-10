package kata

// https://www.codewars.com/kata/526c7363236867513f0005ca/train/go

func IsLeapYear(year int) bool {
	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		return true
	}

	return false
}
