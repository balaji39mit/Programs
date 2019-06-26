package hackerrank

import (
	"github.com/balaji39mit/Programs/DataStructures"
)

var stack DataStructures.IStack

func isBracketsMatching(openingBraces interface{}, closingBraces interface{}) bool {
	if openingBraces == '{' && closingBraces == '}' {
		return true
	}
	if openingBraces == '(' && closingBraces == ')' {
		return true
	}
	if openingBraces == '[' && closingBraces == ']' {
		return true
	}
	return false
}

func isBalanced(expression string) string {
	stack = &DataStructures.StackNode{}
	for _, value := range expression {
		if value == '{' || value == '[' || value == '(' {
			stack.Push(value)
		} else {
			data, err := stack.Pop()
			if err != nil {
				return "NO"
			}
			if !isBracketsMatching(data, value) {
				return "NO"
			}
		}
	}
	if stack.IsEmpty() {
		return "YES"
	}
	return "NO"
}
