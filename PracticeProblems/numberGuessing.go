package practice_problems

import (
	"fmt"
	"strings"
)

func main() {
	var testCase int
	fmt.Scanf("%d", &testCase)

	for i := 0; i < testCase; i++ {
		var a, b, n int
		var text string
		fmt.Scanf("%d %d %d", &a, &b, &n)
		a = a + 1
		for {
			number := (a + b) / 2
			fmt.Printf("%d\n", number)
			fmt.Scanf("%s", &text)
			if strings.EqualFold(text, "CORRECT") {
				break
			} else if strings.EqualFold(text, "TOO_SMALL") {
				a = number + 1
			} else if strings.EqualFold(text, "TOO_BIG") {
				b = number - 1
			} else {
				break
			}
		}
	}
}
