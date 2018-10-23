package main

import (
	"fmt"
)

func merge(left []int32, right []int32) []int32 {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int32, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func mergeSort(arr []int32) []int32 {

	if len(arr) < 2 {
		return arr
	}
	mid := (len(arr)) / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}
func cutSlice(length []int32) []int32 {
	length = mergeSort(length)
	var output []int32
	for len(length) > 0 {
		output = append(output, int32(len(length)))
		cutElement := length[0]
		cutLength := 1
		for i := 1; i < len(length); i++ {
			if length[i] == cutElement {
				cutLength++
			}
			length[i] = length[i] - cutElement
		}
		length = length[cutLength:]
	}
	return output
}

func main() {
	arr := cutSlice([]int32{1, 2, 3, 4, 3, 3, 2, 1})
	fmt.Println(arr)
}
