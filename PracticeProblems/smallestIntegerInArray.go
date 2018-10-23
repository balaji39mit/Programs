package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution1(A []int) int {
	// write your code in Go 1.4
	len := len(A)
	count := make([]int, len+2)
	for i := 0; i < len; i++ {
		if A[i] <= len && A[i] > 0 {
			count[A[i]]++
		}
	}
	for i := 1; i < len+2; i++ {
		if count[i] == 0 {
			return i
		}
	}
	return -1
}

func main() {
	output := Solution1([]int{0, 2, 1})
	fmt.Println(output)
}
