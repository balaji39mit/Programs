package sorting

//BubbleSort sorts the array by doing repeated swaps with the adjacent elements.
//If the adjacent element is smaller than current element, just swap.
//Idea-logy : swap the elements as and when you encounter the mismatch.
//Termination condition : We will be swapping the elements whenever the mismatch occurs. So, the bubble sort ends when swapping does not occur on an iteration.
// How it works? : Let's say in an unsorted array, you are swapping the element since there is a mis-match. But, you are not tracking of the previous indices.
// so, you need to start from the first index on each iteration. Since the array is having n elements and utmost n^2 swaps can occur
// an array will be sorted after n^2 iterations at the worst case.
//After i iterations, last i elements will be placed in their right position.
type BubbleSort struct {
	BaseCase func([]int32) (bool, int)
}

//The typical implementation of bubble sort is stable
// as the key comparison operator swaps the value of two indices which are compared.
// i.e. Unlike Selection Sort, Bubble sort actually swaps the elements which took part in comparison operator.
//So, there is no chance of loosing the relative order here.
func (b *BubbleSort) Sort(arr []int32) []int32 {
	isBaseCase, len := b.BaseCase(arr)
	if isBaseCase {
		return arr
	}
	for i := 0; i < len-1; i++ {
		swapped := false
		//After i iterations, last i elements will be in the proper order.
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}

func (b *BubbleSort) RecursiveSort(arr []int32) []int32 {
	rec(arr, len(arr))
	return arr
}

func rec(arr []int32, n int) {
	if n == 1 {
		return
	}
	//fix the last element.
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			swap(&arr[i], &arr[i+1])
		}
	}
	//Recur for remaining sub-array.
	rec(arr, n-1) // Note : This is known as tail recursion, instead you can point it to start of the function using goto.
}
