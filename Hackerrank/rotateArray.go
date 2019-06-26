package hackerrank

import "fmt"

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	if int32(len(a)) == d {
		return a
	}
	slice := a[0:d]
	fmt.Println("Slice : %v", slice)
	newArray := append(a, slice...)
	fmt.Println("NewArray : %v", newArray)
	rotatedArray := newArray[d:]
	fmt.Println("Rotated : %v", rotatedArray)
	return rotatedArray
}

func rot(a []int32, d int32) []int32 {
	if int32(len(a)) == d {
		return a
	}
	slice := a[0:d]
	fmt.Println("Slice : %v", slice)
	secondSlice := a[d:]
	rotated := append(secondSlice, slice...)
	fmt.Println("Rotated : %v", rotated)
	return rotated
}
