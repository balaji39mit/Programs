package practice_problems

import "fmt"

func factTailRecursion(n uint32, factorial uint32) uint32 {
	if n == 0 {
		return factorial
	}
	return factTailRecursion(n-1, n*factorial)
}

func fact(n uint32) uint32 {
	return factTailRecursion(n, 1)
}

func main() {
	fmt.Println(fact(5))
}
