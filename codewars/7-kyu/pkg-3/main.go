package kata

// 7 kyu
// https://www.codewars.com/kata/58941fec8afa3618c9000184/train/go

func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
	var days, heightPlant int

	for heightPlant < desiredHeight {
		heightPlant += upSpeed

		if heightPlant >= desiredHeight {
			days++
			return days
		}

		heightPlant -= downSpeed
		days++
	}

	return days
}
