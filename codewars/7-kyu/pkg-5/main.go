package kata

// 7 kyu
// https://www.codewars.com/kata/5ba38ba180824a86850000f7/train/go

func Solve(arr []int) []int {
	hash := make(map[int]int, 10)
	res := []int{}

	for i := 0; i < 10; i++ {
		hash[i] = 0
	}

	for _, val := range arr {
		hash[val]++
	}

	for _, val := range arr {
		if hash[val] == 1 || hash[val] == 0 {
			res = append(res, val)
		} else {
			hash[val]--
		}
	}

	return arr

}
