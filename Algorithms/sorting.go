package main

import (
	"fmt"
	"github.com/balaji39mit/Programs/Algorithms/sorting"
)

var (
	sortingTypes map[string]sorting.ISort
)

const (
	MergeSort     = "MergeSort"
	InsertionSort = "InsertionSort"
	SelectionSort = "SelectionSort"
)

func init() {
	sortingTypes[MergeSort] = &sorting.MergeSort{}
	sortingTypes[InsertionSort] = &sorting.InsertionSort{}
	sortingTypes[SelectionSort] = &sorting.SelectionSort{BaseCase: baseCase}
}

var baseCase = func(arr []int32) (isBaseCase bool, n int) {
	n = len(arr)
	isBaseCase = n <= 1
	return
}

func main() {
	//enter the type of sort you want to perform.s
	var arr []int32
	if value, ok := sortingTypes[MergeSort]; ok {
		value.Sort(arr)
	} else {
		fmt.Println("Given sorting type is not supported.")
	}
}
