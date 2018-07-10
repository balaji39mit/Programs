package main

func removeDuplicates(a []int) int {
	len := len(a)
	if len <= 1 {
		return len
	}
	currentElement := a[0]
	nextIndex := 1
	count := 1
	for i := 1; i < len; i++ {
		if a[i] == currentElement {
			continue
		}
		a[nextIndex] = a[i]
		nextIndex++
		currentElement = a[i]
		count++
	}
	return count
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3}
	len := removeDuplicates(arr)
	for i := 0; i < len; i++ {
		print(arr[i])
	}
}
