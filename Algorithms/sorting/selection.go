package sorting

//Time complexity : O(n^2)
//Space complexity : O(1) - In-place sorting
//This typical implementation is not a stable sort.
type SelectionSort struct {
	BaseCase func([]int32) (bool, int)
}

//Helper functions...
func swap(x *int32, y *int32) {
	temp := *x
	*x = *y
	*y = temp
}

//Select the minimum element in each iteration and swap the minimum element with the current element.
//After i iterations, arr[0] to arr[i-1] are sorted and arr[i] to arr[n-1] remains unsorted.
func (s *SelectionSort) Sort(arr []int32) []int32 {
	isBaseCase, length := s.BaseCase(arr)
	if isBaseCase {
		return arr
	}
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			swap(&arr[i], &arr[minIndex])
		}
	}
	return arr
}

//example : 4 2 3 5 4 1
func (s *SelectionSort) StableSort(arr []int32) []int32 {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		key := arr[minIndex]
		for ; minIndex > i; minIndex-- {
			arr[minIndex] = arr[minIndex-1]
		}
		arr[i] = key
	}
	return arr
}
