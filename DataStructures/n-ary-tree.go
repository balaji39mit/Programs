package DataStructures

import "fmt"

type NaryTree struct {
	Val      int
	Children []*NaryTree
}

type INaryTree interface {
	Insert(int)
	Display()
}

func CreateEmptyNaryNode(val int, children int) *NaryTree {
	tree := &NaryTree{
		Val: val,
	}
	if children != 0 {
		tree.Children = make([]*NaryTree, children)
	}
	return tree
}

func (n *NaryTree) Insert(val int, children int) error {
	if n.Val == -1 {
		n.Val = val
		if children != 0 {
			n.Children = make([]*NaryTree, children)
		}
		return nil
	}

	//first level
	for _, item := range n.Children {
		if item.Val == -1 {
			return item.Insert(val, children)
		}
	}

	return nil
}

func (n *NaryTree) PreOrderTraversal() {
	if n == nil || n.Val == -1 {
		return
	}
	fmt.Println("Value : ", n.Val)
	for _, item := range n.Children {
		item.PreOrderTraversal()
	}
}

func (n *NaryTree) MaximumProfit(profitSoFar int) int {
	if n.Val == -1 {
		return profitSoFar
	}
	profitSoFar += n.Val
	maximumValue := -1
	maximumIndex := -1
	for index, item := range n.Children {
		if item.Val > maximumValue {
			maximumIndex = index
			maximumValue = item.Val
		}
	}
	if maximumIndex != -1 {
		return n.Children[maximumIndex].MaximumProfit(profitSoFar)
	}
	return profitSoFar
}

func (n *NaryTree) Display() {
	if n == nil {
		return
	}
	for i := len(n.Children); i > 0; i-- {

	}

}

func init() {

}
