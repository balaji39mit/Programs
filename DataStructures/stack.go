package DataStructures

import "errors"

type StackNode struct {
	data interface{}
	next *StackNode
}

type IStack interface {
	IsEmpty() bool
	Peek() (interface{}, error)
	Push(data interface{})
	Pop() (interface{}, error)
}

func (stack *StackNode) IsEmpty() bool {
	top := stack.next
	return top == nil
}

func (stack *StackNode) Peek() (interface{}, error) {
	top := stack.next
	if top.IsEmpty() {
		return nil, errors.New("Stack is Empty")
	}
	return top.data, nil
}

func (stack *StackNode) Push(data interface{}) {
	top := stack.next
	newNode := &StackNode{data: data}
	newNode.next = top
	stack.next = newNode
}

func (stack *StackNode) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is Empty")
	}
	top := stack.next
	data := top.data
	stack.next = top.next
	return data, nil
}

func init() {

}
