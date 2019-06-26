package main

import (
	"bufio"
	"io"
	"strings"

	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

/*
 * Complete the 'delta_encode' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY numbers as parameter.
 */

func delta_encode(numbers []int32) []int32 {
	n := len(numbers)
	if n <= 0 {
		return []int32{}
	}
	// Write your code here
	output := make([]int32, 0)
	output = append(output, numbers[0])
	for i := 1; i < n; i++ {
		diff := numbers[i] - numbers[i-1]
		if diff < -127 || diff > 127 {
			output = append(output, diff)
		}
		output = append(output, diff)
	}
	return output
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

func _main() {
	numbers := []int32{25626, 25667, 12, 199}
	res := delta_encode(numbers)
	for _, val := range res {
		fmt.Println(val)
	}
}
