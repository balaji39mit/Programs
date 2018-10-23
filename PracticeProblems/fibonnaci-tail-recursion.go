package main

import "fmt"

func FibonacciTR(n uint32, fibonacci uint32) uint32 {
	if n == 0 || n == 1 {
		return fibonacci
	}
	mid := FibonacciTR(n-1, fibonacci)
	return FibonacciTR(n-2, fibonacci+mid)
}

func main() {
	fmt.Println(FibonacciTR(5, 1))
}
