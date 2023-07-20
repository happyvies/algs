package kata

// 8 kyu
// https://www.codewars.com/kata/568dcc3c7f12767a62000038/train/go

func SetAlarm(employed, vacation bool) bool {

	if vacation {
		return false
	}

	return employed || vacation
}
