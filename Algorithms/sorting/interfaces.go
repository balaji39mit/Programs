package sorting

type ISort interface {
	Sort([]int32) []int32
	StableSort([]int32) []int32
}
