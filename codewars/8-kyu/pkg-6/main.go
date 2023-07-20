package kata

// 8 kyu
// https://www.codewars.com/kata/588417e576933b0ec9000045/train/go

func seatsInTheater(nCols int, nRows int, col int, row int) int {
	var total, tRows, tCols int

	total = nCols * nRows
	tRows = row * nCols
	tCols = (col-1)*nRows - row*(col-1)

	return total - (tRows + tCols)
}
