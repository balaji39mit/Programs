package sorting

//Empty struct just to use the ISort interface..s
type MergeSort struct{}

//utility function that merges two sorted input arrays and returns the combined sorted output array
func merge(arr []int32, low, middle, high int) []int32 {
	firstHalf := [middle - low + 1]int32{}
	secondHalf := [high - middle]int32{}
	for i := 0; i < len(firstHalf); i++ {
		firstHalf[i] = arr[low+i]
	}
	for i := 0; i < len(secondHalf); i++ {
		secondHalf[i] = arr[middle+1+i]
	}
	i, j, k := 0, 0, low
	for ; i < len(firstHalf) && j < len(secondHalf); k++ {
		if firstHalf[i] <= secondHalf[j] {
			arr[k] = firstHalf[i]
			i++
		} else {
			arr[k] = secondHalf[j]
			j++
		}
	}
	for ; i < len(firstHalf); i, k = i+1, k+1 {
		arr[k] = firstHalf[i]
	}
	for ; j < len(secondHalf); j, k = j+1, k+1 {
		arr[k] = secondHalf[j]
	}
	return arr
}

//mergeSort sorts the array by recursively calling itself.
func mergeSort(arr []int32, low int, high int) []int32 {
	if low >= high {
		return arr
	}
	//recursively merge the two halves and merge.
	middle := (low) + (high-low)/2
	mergeSort(arr, low, middle)
	mergeSort(arr, middle+1, high)
	return merge(arr, low, middle, high)
}

//Implements the ISort interface using Merge Sort algorithm..
func (s *MergeSort) Sort(arr []int32) []int32 {
	return mergeSort(arr, 0, len(arr)-1)
	return arr
}

func (s *MergeSort) StableSort(arr []int32) []int32 {
	return s.Sort(arr)
}
