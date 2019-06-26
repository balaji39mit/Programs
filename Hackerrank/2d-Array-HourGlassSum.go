package hackerrank

import (
	"bufio"
	"io"
	"math"
	"strings"
)

func max(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var maximumSum int32
	maximumSum = math.MinInt32
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			var sum int32
			sum = 0
			for i := row; i < row+3; i++ {
				for j := col; j < col+3; j++ {
					if (i == row+1) && (j == col || j == col+2) {
						continue
					}
					sum += arr[i][j]
				}
			}
			maximumSum = max(maximumSum, sum)
		}
	}
	return maximumSum
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
