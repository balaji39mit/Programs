package hackerrank

func getMinimumUniqueSum(arr []int32) int32 {
	// Write your code here
	arr = mergeSort(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == arr[i] {
				arr[j]++
			}
		}
	}
	sum := int32(0)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
