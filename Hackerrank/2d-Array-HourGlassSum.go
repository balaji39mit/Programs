package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
