package main

import "fmt"

func countDuplicates(numbers []int32) int32 {
	// Write your code here
	length := len(numbers)
	duplicates := int32(0)
	if length <= 1 {
		return duplicates
	}
	count := make([]int32, 1001)
	for i := 0; i < length; i++ {
		count[numbers[i]]++
	}
	for i := 1; i < len(count); i++ {
		if count[i] > 1 {
			duplicates++
		}
	}
	return duplicates
}

func countDup(a []int32) int32 {

	length := len(a)
	duplicates := int32(0)
	if length <= 1 {
		return duplicates
	}
	count := make(map[int32]int32, 0)
	for i := 0; i < length; i++ {
		if _, ok := count[a[i]]; ok {
			count[a[i]] = count[a[i]] + 1
		} else {
			count[a[i]] = 1
		}
	}

	for _, val := range count {
		if val > 1 {
			duplicates++
		}
	}
	return duplicates
}
func main() {
	fmt.Println(countDuplicates([]int32{}))
}
