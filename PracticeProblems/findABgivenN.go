package practice_problems

import (
	"fmt"
	"strconv"

	"strings"
)

func contains4(x uint64) bool {
	str := fmt.Sprintf("%d", x)
	return strings.Contains(str, "4")
}

func optimalSol(target uint64) (a, b uint64) {
	str := fmt.Sprintf("%d", target)
	arrayStr := strings.Split(str, "")
	for key, value := range arrayStr {
		if value == "4" {
			arrayStr[key] = "2"
		} else {
			arrayStr[key] = "0"
		}
	}
	a, _ = strconv.ParseUint(strings.Join(arrayStr, ""), 10, 64)

	str = fmt.Sprintf("%d", target)
	arrayStr = strings.Split(str, "")
	for key, value := range arrayStr {
		if value == "4" {
			arrayStr[key] = "2"
		}
	}
	b, _ = strconv.ParseUint(strings.Join(arrayStr, ""), 10, 64)
	return
}

func binarySearch(low uint64, high uint64, target uint64) (a, b uint64) {
start:
	if low+high == target {
		if !contains4(low) && !contains4(high) {
			return low, high
		}
	}
	mid := low + (high-low)/2
	low = low + (mid-low)/2
	if mid%2 != 0 {
		low += 1
	}
	high = mid + (high-mid)/2
	if low > high {
		return 0, 0
	}
	goto start
}
