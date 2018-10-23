package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*

Problem statement : Given an array of integers which represents a number. Add 1 to the number and return the number in array form.

Input : [1, 2, 3, 4] : 1234 is a number

Output : [1. 2. 3. 5] : After adding 1 to a number 1234, we will get the result as 1235.

Notes :

0 <= len(array) <= 1000

0 <= array[i] <= 9

If the above is solved, modify the constraint with the input range as -9 <= array[i] <= 9
*/

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

func reverseArray(input []int32) []int32 {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func subtractOne(arr []int32) []int32 {
	carry := int32(-1)
	var output []int32
	for i := len(arr) - 1; i >= 0; i-- {
		value := arr[i]
		if arr[i] == 0 {
			value = 10
		}
		value = value + carry
		output = append(output, value)
		//set the carry since the current number borrowed one from previous number
		if arr[i] == 0 {
			carry = -1
		} else {
			carry = 0
		}
	}
	result := reverseArray(output)
	if result[0] == 0 && len(arr) > 1 {
		result = result[1:]
	}
	result[0] = result[0] * -1
	return result
}

func addOne(arr []int32) []int32 {
	carry := int32(1)
	var output []int32
	if len(arr) == 0 {
		return []int32{1}
	}
	if arr[0] < 0 {
		arr[0] = -1 * arr[0]
		return subtractOne(arr)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		value := arr[i] + carry
		carry = value / 10
		value = value % 10
		output = append(output, value)
	}
	if carry > 0 {
		output = append(output, carry)
	}
	return reverseArray(output)
}

func main() {

	fmt.Println("Add 1 to 1234 : %v", addOne([]int32{1, 2, 3, 4}))
	fmt.Println("Add 1 to empty array : %v", addOne([]int32{}))
	fmt.Println("Add 1 to 99 : %v", addOne([]int32{9, 9}))
	fmt.Println("Add 1 to 1000 : %v", addOne([]int32{1, 0, 0, 0}))

	fmt.Println("Negative integers....\n")
	fmt.Println("Add 1 to -1 : %v", addOne([]int32{-1}))
	fmt.Println("Add 1 to -8 : %v", addOne([]int32{-8}))
	fmt.Println("Add 1 to -110 : %v", addOne([]int32{-1, 1, 0}))
	fmt.Println("Add 1 to -100 : %v", addOne([]int32{-1, 0, 0}))

	/*reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriter(os.Stdout)
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	n := int32(nTemp)
	checkError(err)
	var arr []int32
	arrTemp := strings.Split(readLine(reader), " ")
	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, int32(arrItem))
	}

	output := addOne(arr)

	fmt.Fprintf(writer, "%v", output)

	writer.Flush()*/

}
