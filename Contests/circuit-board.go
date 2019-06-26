package main

type BoardIndex struct {
	Min    int
	Max    int
	Length int
}

func NumberOfSquares(matrix [][]int64, k int64) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	if cols == 1 {
		return rows
	}
	goodBoard := make(map[int][]BoardIndex, 0)
	for key, val := range matrix {
		//traversing the row.
		minIndex := 0
		maxIndex := 0
		for i := 1; i < len(val); i++ {
			if matrix[key][i]-matrix[key][minIndex] > k {
				goodBoard[key] = append(goodBoard[key], BoardIndex{minIndex, maxIndex, maxIndex - minIndex + 1})
				minIndex = i
			}
			maxIndex = i
		}
		goodBoard[key] = append(goodBoard[key], BoardIndex{minIndex, maxIndex, maxIndex - minIndex + 1})
	}
	//Now goodBoard has all the indexes of good sub circuit board. Try to find the maximum sub-rectangle out of these indexes.
	for key, val := range goodBoard {

	}
	return 0
}
