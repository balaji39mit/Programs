package main

import "github.com/balaji39mit/Programs/DataStructures"

var sortedList1, sortedList2 *DataStructures.ListNode

func mergeTwoSortedList(list1 *DataStructures.ListNode, list2 *DataStructures.ListNode) *DataStructures.ListNode {
	sortedList := &DataStructures.ListNode{}
	traversalNode := sortedList
	for ; list1 != nil && list2 != nil; traversalNode = traversalNode.Next {
		if list1.Data.(int) <= list2.Data.(int) {
			newNode := &DataStructures.ListNode{Data: list1.Data}
			traversalNode.Next = newNode
			list1 = list1.Next
		} else {
			newNode := &DataStructures.ListNode{Data: list2.Data}
			traversalNode.Next = newNode
			list2 = list2.Next
		}
	}
	for ; list1 != nil; list1, traversalNode = list1.Next, traversalNode.Next {
		newNode := &DataStructures.ListNode{Data: list1.Data}
		traversalNode.Next = newNode
	}
	for ; list2 != nil; list2, traversalNode = list2.Next, traversalNode.Next {
		newNode := &DataStructures.ListNode{Data: list2.Data}
		traversalNode.Next = newNode
	}
	return sortedList
}

func main() {
	sortedList1 = &DataStructures.ListNode{}
	sortedList2 = &DataStructures.ListNode{}
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
