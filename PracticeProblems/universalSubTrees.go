package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

//returns the count of non-empty universal sub trees and
//whether the given sub-tree is universal or not
func countUniversalSubTree(root *Node) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftCount, leftOk := countUniversalSubTree(root.left)
	rightCount, rightOk := countUniversalSubTree(root.right)
	totalCount := leftCount + rightCount
	if ok := leftOk && rightOk; ok {
		if root.left != nil && root.left.value != root.value {
			return totalCount, false
		}
		if root.right != nil && root.right.value != root.value {
			return totalCount, false
		}
		return totalCount + 1, true
	}
	return totalCount, false
}

func main() {
	fmt.Println(countUniversalSubTree(&Node{}))
}
