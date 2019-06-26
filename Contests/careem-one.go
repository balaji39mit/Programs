package main

import (
	"sort"

	"fmt"
)

//Custom map to sort based on value
type SortMap struct {
	key int
	val int32
}

func widestGap(n int32, start []int32, finish []int32) int32 {
	// Write your code here

	//Populating the rear of each car to a custom map which will be used for sorting the rear of the car
	sortedMap := make([]SortMap, len(start))
	for i := 0; i < len(start); i++ {
		sortedMap[i] = SortMap{i, start[i]}
	}

	//Sort the map based on value in ascending order.
	sort.Slice(sortedMap, func(i, j int) bool {
		return sortedMap[i].val < sortedMap[j].val
	})

	endSoFar := int32(0) //Used for the maximum front of the car at any time.
	res := int32(0)      // Used for fining the largest gap at any point in time.

	//Logic : Traverse through the sortedMap which is sorted in ascending order based on the rear of the car.
	//For each index, find the difference with the endSoFar variable and track the maximum.
	for i := 1; i < len(sortedMap); i++ {
		key := sortedMap[i-1].key
		if finish[key] > endSoFar {
			endSoFar = finish[key]
		}
		diff := sortedMap[i].val - endSoFar - 1
		if diff > res {
			res = diff
		}
	}

	//The last element's front position is missed in the above loop
	//So, update endSoFar variable at this stage and maintain the definition of the variable
	m := len(start)
	lastElem := finish[sortedMap[m-1].key]
	if lastElem > endSoFar {
		endSoFar = lastElem
	}

	//There can be cases where gaps are not present between any of the cars.
	//So find whether any gap is there for the last car with the length of the road.
	diff := n - endSoFar
	if diff > res {
		res = diff
	}

	return res
}

func main() {
	/*start := []int32{22,
		75,
		26,
		45,
		72,
		81,
		47,
		29,
		97,
		2,
		75,
		25,
		82,
		84,
		17,
		56,
		32,
		2,
		28,
		37,
		57,
		39,
		18,
		11,
		79,
		6,
		40,
		68,
		68,
		16,
		40,
		63,
		93,
		49,
		91,
		10,
		55,
		68,
		31,
		80}
	finish := []int32{51,
		92,
		59,
		60,
		77,
		95,
		61,
		68,
		98,
		90,
		87,
		39,
		94,
		85,
		67,
		74,
		41,
		65,
		78,
		80,
		85,
		93,
		87,
		82,
		83,
		16,
		89,
		81,
		69,
		72,
		80,
		77,
		99,
		90,
		92,
		95,
		68,
		70,
		75,
		97}
	fmt.Println(widestGap(100, start, finish))*/
	start := []int32{3, 8}
	end := []int32{4, 9}
	fmt.Println(widestGap(10, start, end))

}
