package Coursera

func mergeAndCountSplitInv(arr []int32, low, mid, high int) int {
	firstHalf := [mid - low + 1]int32{}
	secondHalf := [high - mid]int32{}
	for i := 0; i < len(firstHalf); i++ {
		firstHalf[i] = arr[low+i]
	}
	for i := 0; i < len(secondHalf); i++ {
		secondHalf[i] = arr[mid+1+i]
	}
	i, j, k, count := 0, 0, low, 0
	for ; i < len(firstHalf) && j < len(secondHalf); k++ {
		if firstHalf[i] <= secondHalf[j] {
			arr[k] = firstHalf[i]
			i++
		} else {
			arr[k] = secondHalf[j]
			count += len(firstHalf) - i
			j++
		}
	}
	for ; i < len(firstHalf); i, k = i+1, k+1 {
		arr[k] = firstHalf[i]
	}
	for ; j < len(secondHalf); j, k = j+1, k+1 {
		arr[k] = secondHalf[j]
	}
	return count
}

func sortAndCount(arr []int32, low, high int) int {
	if low <= high {
		return 0
	}
	mid := low + (high-low)/2
	leftInversion := sortAndCount(arr, low, mid)
	rightInversion := sortAndCount(arr, mid+1, high)
	splitInv := mergeAndCountSplitInv(arr, low, mid, high)
	return leftInversion + rightInversion + splitInv
}

func CountInversion(arr []int32) int {
	return sortAndCount(arr, 0, len(arr)-1)
}
