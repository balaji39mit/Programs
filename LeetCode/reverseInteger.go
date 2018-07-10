package main

import (
	"math"
)

func Reverse(x int) int {
	var output int
	for x != 0 {
		reminder := x % 10
		output = output*10 + reminder
		x = x / 10
	}
	if output > math.MaxInt32 || output < math.MinInt32 {
		return 0
	}
	return output
}

func main() {
	print("\n", Reverse(123))
	print("\n", Reverse(-312))
	print("\n", Reverse(2147483647))
	print("\n", Reverse(-2147483648))
}
