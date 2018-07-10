package main

func twoSumIndices(a []int, target int) []int {
	len := len(a)
	if len <= 1 {
		return []int{}
	}
	visitedSoFar := map[int]int{
		a[0]: 0,
	}
	for i := 1; i < len; i++ {
		if val, ok := visitedSoFar[target-a[i]]; ok {
			return []int{val, i}
		}
		visitedSoFar[a[i]] = i
	}
	return []int{}
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 15
	output := twoSumIndices(arr, target)
	for _, value := range output {
		print(value)
	}
}
