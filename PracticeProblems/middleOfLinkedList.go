package practice_problems

import (
	"github.com/balaji39mit/Programs/DataStructures"
)

var linkedList *DataStructures.LinkedList

func init() {
	linkedList = &DataStructures.LinkedList{
		Head: &DataStructures.ListNode{
			Data: 1,
			Next: &DataStructures.ListNode{
				Data: 2,
				Next: &DataStructures.ListNode{
					Data: 3,
					Next: &DataStructures.ListNode{
						Data: 4,
						Next: &DataStructures.ListNode{
							Data: 5,
						},
					},
				},
			},
		},
	}
}

func GetMiddleElement(list *DataStructures.LinkedList) interface{} {
	if list == nil || list.Head == nil {
		return nil
	}
	slow, fast := list.Head, list.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Data
}
