package practice_problems

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution2(S string) int {
	// write your code in Go 1.4
	len := len(S)
	count := 1 // For 0th Cyclic.
	for k := 1; k < len; k++ {
		//S[low:] -> fetches all the character from low till end of the string.
		//S[:high] -> fetch from the first character till the (high-1)th position of the string. high is excluded here.
		tempStr := S[k:] + S[:k]
		//fmt.Println(tempStr)
		if tempStr == S {
			count++
		}
	}
	return count
}
