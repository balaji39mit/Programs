package main

import "fmt"

// Complete the maxDifference function below.
func maxDifference(a []int32) int32 {
	len := len(a)
	maxDiff := int32(-1)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if a[i] < a[j] && maxDiff < (a[j]-a[i]) {
				maxDiff = a[j] - a[i]
			}
		}
	}
	return maxDiff
}

func optimizedMaxDifference(a []int32) int32 {
	len := len(a)
	maxDiff := int32(-1)
	if len <= 1 {
		return maxDiff
	}
	minimumElement := a[0]
	for i := 1; i < len; i++ {
		if a[i] > minimumElement && maxDiff < (a[i]-minimumElement) {
			maxDiff = a[i] - minimumElement
		} else if a[i] < minimumElement {
			minimumElement = a[i]
		}
	}
	return maxDiff
}

func main() {
	fmt.Println(maxDifference([]int32{}))
}
