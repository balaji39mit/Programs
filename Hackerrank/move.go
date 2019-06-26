//Problem statement : GGiven an array of integers with 0'a and 1's. Move the elements such that 0's should be one side and 1's should be on the other side.
//Rules are : 1. you can move only adjacent elements.
//2. The number of moves should be minimum.
//3. Return the minimum number of moves.

package hackerrank

func process(arr []int32) bool {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			count++
		}
	}
	return count <= 1
}

func moveZerotoRight(arr []int32, count int32) int32 {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && arr[i+1] == 1 {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			count++
		}
	}
	if process(arr) {
		return count
	}
	return moveZerotoRight(arr, count)
}

func moveZerotoLeft(arr []int32, count int32) int32 {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 && arr[i+1] == 0 {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			count++
		}
	}
	if process(arr) {
		return count
	}
	return moveZerotoLeft(arr, count)
}

func minMoves(avg []int32) int32 {
	if process(avg) {
		return 0
	}
	// Write your code here
	min1 := moveZerotoLeft(avg, 0)
	min2 := moveZerotoRight(avg, 0)
	if min1 <= min2 {
		return min1
	}
	return min2
}
