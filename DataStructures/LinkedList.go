package DataStructures

import (
	"errors"
	"fmt"
)

type ListNode struct {
	Data interface{}
	Next *ListNode
}

type ILinkedList interface {
	//Inserts an element into a linked list at the end.
	Insert(data interface{})
	//Inserts an element into a linked list at the beginning.
	InsertBeg(data interface{})
	//Inserts an element into a linked list at the middle position.
	//Returns an error if invalid position has been passed.
	InsertMid(data interface{}, position int) error
	//Display the elements present in the linked list.
	Display()
	//Delete the last element from the linked list
	//and return the deleted element.
	Delete() (data interface{}, err error)
	//Delete the first element.
	//and return the deleted element.
	DeleteBeg() (data interface{}, err error)
	//Delete the middle element.
	//and return the deleted element.
	//Error is returned if invalid position has been passed.
	DeleteMid(position int) (data interface{}, err error)
	//Returns the length of the linked list.
	Length() int
	//Tells whether the given data is present or not.
	IsPresent(data interface{}) bool
	//Returns the index of the given element if present.
	Search(data interface{}) (int, error)
	//Reverse the given linked list.
	Reverse()
}

func (list *ListNode) Insert(data interface{}) {
	newNode := &ListNode{Data: data}
	temp := list
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode
}

func (list *ListNode) InsertBeg(data interface{}) {
	newNode := &ListNode{Data: data}
	currentNode := list.Next
	list.Next = newNode
	newNode.Next = currentNode
}

func (list *ListNode) InsertMid(data interface{}, position int) error {
	if position > list.Length() || position < 0 {
		return errors.New("Invalid position.")
	}
	temp := list
	for count := 0; count < position; count++ {
		temp = temp.Next
	}
	nextNode := temp.Next
	newNode := &ListNode{Data: data}
	temp.Next = newNode
	newNode.Next = nextNode
	return nil
}

func (list *ListNode) Display() {
	for temp := list.Next; temp != nil; temp = temp.Next {
		print(fmt.Sprintf("%v ", temp.Data))
	}
}

func (list *ListNode) Delete() (interface{}, error) {
	return nil, nil
}

func (list *ListNode) DeleteBeg() (interface{}, error) {
	return nil, nil
}

func (list *ListNode) DeleteMid(position int) (interface{}, error) {
	return nil, nil
}

func (list *ListNode) Length() int {
	count := 0
	for temp := list.Next; temp != nil; temp = temp.Next {
		count++
	}
	return count
}

func (list *ListNode) IsPresent(data interface{}) bool {
	for temp := list.Next; temp != nil; temp = temp.Next {
		if temp.Data == data {
			return true
		}
	}
	return false
}

func (list *ListNode) Search(data interface{}) int {
	count := 0
	for temp := list.Next; temp != nil; temp, count = temp.Next, count+1 {
		if temp.Data == data {
			return count
		}
	}
	return -1
}

func (list *ListNode) Reverse() {

}

func init() {

}
