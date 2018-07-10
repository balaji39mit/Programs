package main

func removeElement(nums []int, value int) int {
	currentIndex := 0
	for index, val := range nums {
		if val != value {
			nums[currentIndex] = nums[index]
			currentIndex = currentIndex + 1
		}
	}
	return currentIndex
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	value := 2
	len := removeElement(nums, value)
	for i := 0; i < len; i++ {
		print(nums[i])
	}

}
