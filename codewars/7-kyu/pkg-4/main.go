package kata

// 7 kyu
// https://www.codewars.com/kata/62c0ad702af4fc0023ad5b89/train/go

func PickGrains(grains <-chan string) (good int, bad int) {
	// TODO: your implementation goes here
	for val := range grains {
		if val == "good" {
			good++
		} else {
			bad++
		}
	}
	return
}
