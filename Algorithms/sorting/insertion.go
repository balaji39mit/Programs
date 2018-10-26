package sorting

//Insertion sort works in the way of sorting the cards in hands.
//i.e. You will take a card, you will just compare it with all the elements to the left of the card
//i.e. and insert the current card in the correct position, so that all the cards which are positioned at left
//are moved one step forward to the right.
//Insertion sort is in-place sorting algorithm which makes sure that
//after ith iteration, a[0..i-1] are sorted. i.e. i elements are sorted.
//Works well in best case, since there is no need to swap any element. 0(n)
//worst case occurs when array is sorted reversely, i.e. 0(n^2)
type InsertionSort struct{}

func (s *InsertionSort) Sort(arr []int32) []int32 {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	for i := 1; i < len; i++ {
		j := i - 1
		for j >= 0 && arr[j] > arr[i] {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = arr[i]
	}
	return arr
}

func (s *InsertionSort) StableSort(arr []int32) []int32 {
	return s.Sort(arr)
}

func recursive(arr []int32, index int) []int32 {
	if index >= len(arr) {
		return arr
	}
	low := index - 1
	for low >= 0 && arr[low] > arr[index] {
		arr[low+1] = arr[low]
		low = low - 1
	}
	arr[low+1] = arr[index]
	return recursive(arr, index+1) // tail recursion, can be replaced with goto.
}

func (s *InsertionSort) Recursive(arr []int32) []int32 {
	return recursive(arr, 1)
}
