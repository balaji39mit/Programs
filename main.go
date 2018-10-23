package main

import (
	"fmt"
	"github.com/balaji39mit/Programs/DataStructures"
)

func binaryTree() {
	root := DataStructures.CreateEmptyNode()
	fmt.Println("Enter the number of nodes in the tree.")
	var n int
	fmt.Scanln(&n)
	var x int
	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		fmt.Println("Value of x is : %v", x)
		root.Insert(x)
	}
	fmt.Println("Displaying the elements present in the tree.")
	root.Display(0)
}

func NaryTree() {
	/*root := DataStructures.CreateEmptyNaryNode(-1, 0)
	fmt.Println("Enter the number of nodes in the tree.")
	var n int
	fmt.Scanln(&n)
	var value, children int
	for i := 0; i < n; i++ {
		fmt.Println("Enter the value : ")
		fmt.Scan(&value)
		fmt.Println("Enter the number of children for this node (Note : Please enter 0 for leaf nodes) : ")
		fmt.Scan(&children)
		root.Insert(value, children)
	}*/
	//For now, just hard-code the nary tree.

	root := DataStructures.CreateEmptyNaryNode(12, 3)
	root.Children[0] = DataStructures.CreateEmptyNaryNode(9, 1)
	root.Children[0].Children[0] = DataStructures.CreateEmptyNaryNode(13, 0)
	root.Children[1] = DataStructures.CreateEmptyNaryNode(25, 0)
	root.Children[2] = DataStructures.CreateEmptyNaryNode(16, 2)
	root.Children[2].Children[0] = DataStructures.CreateEmptyNaryNode(13, 0)
	root.Children[2].Children[1] = DataStructures.CreateEmptyNaryNode(9, 0)
	root.PreOrderTraversal()
	fmt.Println("Maximum profit is : ", root.MaximumProfit())
}

func main() {
	//binaryTree()
	NaryTree()
}
