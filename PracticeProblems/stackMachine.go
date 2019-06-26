package practice_problems

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

//hard-coded strings, constants and variable declarations
var (
	duplicate       = "DUP"
	pop             = "POP"
	addition        = "+"
	subtraction     = "-"
	stringSeparator = " "
	MININT          = uint32(0)
	MAXINT          = uint32((math.Pow(2, 20)) - 1)
	stackMachine    IStack
)

//Implementation of stack.
type StackNode struct {
	data uint32
	next *StackNode
}

type IStack interface {
	IsEmpty() bool
	Top() (uint32, error)
	Push(data uint32) error
	Pop() (uint32, error)
	performDuplicate() bool
	performPop() bool
	performAddition() bool
	performSubtraction() bool
	performPush(string) bool
}

//Returns true if the stack is empty..
func (stack *StackNode) IsEmpty() bool {
	top := stack.next
	return top == nil
}

//Returns the top most element..
func (stack *StackNode) Top() (uint32, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is Empty")
	}
	top := stack.next
	return top.data, nil
}

//Insert a new element into the stack..
//Additionally, checks for Overlfow exceptions. Allowed range is 0 <= (2 ^ 20) - 1
func (stack *StackNode) Push(data uint32) error {
	if data < MININT || data > MAXINT {
		return errors.New("Overflow occurred")
	}
	top := stack.next
	newNode := &StackNode{data: data}
	newNode.next = top
	stack.next = newNode
	return nil
}

//Deletes the top element from the stack
func (stack *StackNode) Pop() (uint32, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is Empty")
	}
	top := stack.next
	data := top.data
	stack.next = top.next
	return data, nil
}

//FetchStruct the top element from the stack and
//push that element to the stack
func (stack *StackNode) performDuplicate() bool {
	top, err := stack.Top()
	if err != nil {
		//		fmt.Println(err)
		return false
	}
	err = stack.Push(top)
	if err != nil {
		//		fmt.Println(err)
		return false
	}
	return true
}

//Pops the element from the stack and acknowledge the same.
func (stack *StackNode) performPop() bool {
	_, err := stack.Pop()
	if err != nil {
		//	fmt.Println(err)
		return false
	}
	return true
}

//pop the topmost two elements from the stack and
//push the addition of two elements into the stack.
func (stack *StackNode) performAddition() bool {
	data1, err := stack.Pop()
	if err != nil {
		//	fmt.Println(err)
		return false
	}
	data2, err := stack.Pop()
	if err != nil {
		//	fmt.Println(err)
		return false
	}
	err = stack.Push(data1 + data2)
	if err != nil {
		//	fmt.Println(err)
		return false
	}
	return true
}

//pop the topmost two elements from stack and
//push the difference between them into the stack. (i.e. top - secondTop)
func (stack *StackNode) performSubtraction() bool {
	data1, err := stack.Pop()
	if err != nil {
		return false
	}
	data2, err := stack.Pop()
	if err != nil {
		return false
	}
	err = stack.Push(data1 - data2)
	if err != nil {
		return false
	}
	return true
}

//Converts a given string into integer and pushes the int to stack...
func (stack *StackNode) performPush(operation string) bool {
	number, err := strconv.Atoi(operation)
	if err != nil {
		//	fmt.Println(err)
		return false
	}
	err = stack.Push(uint32(number))
	if err != nil {
		//	fmt.Println(err)
		return false
	}
	return true
}

func Solution(S string) int {
	// write your code in Go 1.4
	stackMachine = &StackNode{}
	operations := strings.Split(S, stringSeparator)
	for _, operation := range operations {
		switch operation {
		case duplicate:
			if !stackMachine.performDuplicate() {
				return -1
			}
		case pop:
			if !stackMachine.performPop() {
				return -1
			}
		case addition:
			if !stackMachine.performAddition() {
				return -1
			}
		case subtraction:
			if !stackMachine.performSubtraction() {
				return -1
			}
		default:
			if !stackMachine.performPush(operation) {
				return -1
			}
		}
	}
	top, err := stackMachine.Top()
	if err != nil {
		return -1
	}
	return int(top)
}
