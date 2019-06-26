package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	var k int
	fmt.Scanf("%d", &k)
	res := 0
	hashMap := make(map[string]struct{}, 0)
	for i := 0; i < n; i++ {
		counter := k
		str := ""
		for j := i; j < n; j++ {
			ok := false
			if arr[j]&1 == 0 {
				ok = true
			} else if counter > 0 {
				ok = true
				counter--
			} else {
				break
			}
			str += strconv.Itoa(arr[j])
			if ok {
				if _, ok1 := hashMap[str]; !ok1 {
					hashMap[str] = struct{}{}
					res++
				}
			}
		}
	}
	fmt.Printf("%d", res)

}
