package DataStructures

import (
	"errors"
	"fmt"
)

type ListNode struct {
	Data interface{}
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

type ILinkedList interface {
	//Inserts an element into a linked list at the end.
	Insert(data interface{})
	//Inserts an element into a linked list at the beginning.
	InsertBeg(data interface{})
	//Inserts an element into a linked list at the middle position.
	//Returns an error if invalid position has been passed.
	//position starts from 0 like array.
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
	//Deletes the given element from the linked list if present.
	//Returns error if the given element is not present.
	DeleteElement(data interface{}) (interface{}, error)
	//Returns the length of the linked list.
	Length() int
	//Tells whether the given data is present or not.
	IsPresent(data interface{}) bool
	//Returns the index of the given element if present.
	//Assumption : LinkedList Index starts from 0 like array.
	Search(data interface{}) int
	//Reverse the given linked list.
	Reverse()
}

func (list *LinkedList) Insert(data interface{}) {
	newNode := &ListNode{Data: data}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	temp := list.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode
}

func (list *LinkedList) InsertBeg(data interface{}) {
	newNode := &ListNode{Data: data}
	newNode.Next = list.Head
	list.Head = newNode
}

func (list *LinkedList) InsertMid(data interface{}, position int) error {
	if position > list.Length() || position < 0 {
		return errors.New("invalid position")
	}
	if position == 0 {
		list.InsertBeg(data)
		return nil
	}
	temp := list.Head
	//We don't need to worry about the nil pointer exception, as the length check is already made.
	for count := 0; count < position-1; count++ {
		temp = temp.Next
	}
	nextNode := temp.Next
	newNode := &ListNode{Data: data}
	temp.Next = newNode
	newNode.Next = nextNode
	return nil
}

func (list *LinkedList) Display() {
	for temp := list.Head; temp != nil; temp = temp.Next {
		fmt.Print(fmt.Sprintf("%v ", temp.Data))
	}
}

func (list *LinkedList) Delete() (interface{}, error) {
	if list.Length() == 0 {
		return nil, errors.New("sorry, linked list is empty. can't delete")
	}
	if list.Head.Next == nil {
		data := list.Head.Data
		list.Head = nil
		return data, nil
	}
	temp := list.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	data := temp.Next.Data
	temp.Next = nil
	return data, nil
}

func (list *LinkedList) DeleteBeg() (interface{}, error) {
	if list.Length() == 0 {
		return nil, errors.New("sorry, linked list is empty. can't delete")
	}
	temp := list.Head
	list.Head = temp.Next
	return temp.Data, nil
}

func (list *LinkedList) DeleteMid(position int) (interface{}, error) {
	if position >= list.Length() || position < 0 {
		return nil, errors.New("invalid position")
	}
	if position == 0 {
		return list.DeleteBeg()
	}
	count, temp := 0, list.Head
	for count < position-1 {
		count, temp = count+1, temp.Next
	}
	nextNode := temp.Next
	temp.Next = nextNode.Next
	return nextNode.Data, nil
}

func (list *LinkedList) DeleteElement(data interface{}) (interface{}, error) {
	if index := list.Search(data); index == -1 {
		return nil, errors.New("sorry, given element is not present in linked list")
	} else {
		return list.DeleteMid(index)
	}
}

func (list *LinkedList) Length() int {
	count := 0
	for temp := list.Head; temp != nil; temp = temp.Next {
		count++
	}
	return count
}

func (list *LinkedList) IsPresent(data interface{}) bool {
	for temp := list.Head; temp != nil; temp = temp.Next {
		if temp.Data == data {
			return true
		}
	}
	return false
}

func (list *LinkedList) Search(data interface{}) int {
	count := 0
	for temp := list.Head; temp != nil; temp, count = temp.Next, count+1 {
		if temp.Data == data {
			return count
		}
	}
	return -1
}

func (list *LinkedList) ReverseRecurse() {

}

func (list *LinkedList) Reverse() {
	var prev *ListNode
	current := list.Head
	prev = nil
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	list.Head = prev
}

func (list *LinkedList) IsLoopPresent() (bool, *ListNode) {
	if list.Head == nil {
		return false, nil
	}
	slow, fast := list.Head, list.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true, slow
		}
	}
	return false, nil
}

func (list *LinkedList) FindLoopNode() *ListNode {
	if ok, meetNode := list.IsLoopPresent(); ok {
		start := list.Head
		for start != meetNode {
			start = start.Next
			meetNode = meetNode.Next
		}
		return start
	}
	return nil
}

func init() {

}
