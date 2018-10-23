package sorting

//Empty struct just to use the ISort interface..s
type MergeSort struct{}

//Implements the ISort interface using Merge Sort algorithm..
func (s *MergeSort) Sort(arr []int32) []int32 {
	return arr
}

func (s *MergeSort) StableSort(arr []int32) []int32 {
	return s.Sort(arr)
}
