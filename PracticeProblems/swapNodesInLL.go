package practice_problems

import "github.com/balaji39mit/Programs/DataStructures"

var ll DataStructures.ILinkedList

func init() {
	ll = &DataStructures.LinkedList{}
}

func main() {
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)
	ll.Display()

	ll = swapNodes(ll.(*DataStructures.LinkedList), 2, 4)
	ll.Display()
}

func swapNodes(ll *DataStructures.LinkedList, x, y int64) *DataStructures.LinkedList {

	return ll
}
