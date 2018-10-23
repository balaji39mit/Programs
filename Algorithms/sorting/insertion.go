package sorting

type InsertionSort struct{}

func (s *InsertionSort) Sort(arr []int32) []int32 {
	return arr
}

func (s *InsertionSort) StableSort(arr []int32) []int32 {
	return s.Sort(arr)
}
