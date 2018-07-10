package main

import (
	"github.com/balaji39mit/Programs/DataStructures"
)

var InputList1, InputList2 *DataStructures.ListNode

func addTwo(list1 *DataStructures.ListNode, list2 *DataStructures.ListNode) *DataStructures.ListNode {
	list := &DataStructures.ListNode{}
	carry := 0
	for ; list1 != nil && list2 != nil; list1, list2 = list1.Next, list2.Next {
		data1 := (list1.Data).(int)
		data2 := (list2.Data).(int)
		sum := data1 + data2 + carry
		carry = sum / 10
		sum = sum % 10
		list.Insert(sum)
	}
	for ; list1 != nil; list1 = list1.Next {
		data := (list1.Data).(int)
		sum := data + carry
		carry = sum / 10
		sum = sum % 10
		list.Insert(sum)
	}
	for ; list2 != nil; list2 = list2.Next {
		data := (list2.Data).(int)
		sum := data + carry
		carry = sum / 10
		sum = sum % 10
		list.Insert(sum)
	}
	if carry != 0 {
		list.Insert(carry)
	}
	return list
}

func main() {
	InputList1 = &DataStructures.ListNode{}
	InputList2 = &DataStructures.ListNode{}
	InputList1.Insert(2)
	InputList1.Insert(4)
	InputList1.Insert(3)
	InputList2.Insert(5)
	InputList2.Insert(6)
	InputList2.Insert(4)
	print("Displaying linked list 1\n")
	InputList1.Display()
	print("\nDisplaying linked list 2\n")
	InputList2.Display()
	List := addTwo(InputList1.Next, InputList2.Next)
	print("\nDisplaying output...\n")
	List.Display()
}
