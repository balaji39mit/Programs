package main

import "github.com/balaji39mit/Programs/DataStructures"

var sortedList1, sortedList2 *DataStructures.LinkedList

func mergeTwoSortedList(list1 *DataStructures.LinkedList, list2 *DataStructures.LinkedList) *DataStructures.LinkedList {
	sortedList := &DataStructures.LinkedList{}
	traversalNode := sortedList
	for ; list1 != nil && list2 != nil; traversalNode = traversalNode.Next {
		if list1.Data.(int) <= list2.Data.(int) {
			newNode := &DataStructures.LinkedList{Data: list1.Data}
			traversalNode.Next = newNode
			list1 = list1.Next
		} else {
			newNode := &DataStructures.LinkedList{Data: list2.Data}
			traversalNode.Next = newNode
			list2 = list2.Next
		}
	}
	for ; list1 != nil; list1, traversalNode = list1.Next, traversalNode.Next {
		newNode := &DataStructures.LinkedList{Data: list1.Data}
		traversalNode.Next = newNode
	}
	for ; list2 != nil; list2, traversalNode = list2.Next, traversalNode.Next {
		newNode := &DataStructures.LinkedList{Data: list2.Data}
		traversalNode.Next = newNode
	}
	return sortedList
}

func main() {
	sortedList1 = &DataStructures.LinkedList{}
	sortedList2 = &DataStructures.LinkedList{}
	sortedList1.Insert(1)
	sortedList1.Insert(2)
	sortedList1.Insert(4)
	sortedList2.Insert(1)
	sortedList2.Insert(3)
	sortedList2.Insert(4)
	print("Displaying linked list 1\n")
	sortedList1.Display()
	print("\nDisplaying linked list 2\n")
	sortedList2.Display()
	List := mergeTwoSortedList(sortedList1.Next, sortedList2.Next)
	print("\nDisplaying output...\n")
	List.Display()

}
