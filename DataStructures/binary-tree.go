package DataStructures

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ITreeNode interface {
	Insert(int) *TreeNode
	Display()
	Delete(int) error
	Length() int
	IsPresent(int) bool
	IsEmpty() bool
	IsBST() bool
}

func CreateEmptyNode() *TreeNode {
	return &TreeNode{
		Val: -1,
	}
}

func (t *TreeNode) Insert(val int) error {
	//root is nil
	if t.Val == -1 {
		t.Val = val
		return nil
	}
	if t.Val < val {
		if t.Right == nil {
			t.Right = CreateEmptyNode()
		}
		return t.Right.Insert(val)
	}
	if t.Left == nil {
		t.Left = CreateEmptyNode()
	}
	return t.Left.Insert(val)
}

func (t *TreeNode) Display(depth int) {
	if t.Val == -1 {
		return
	}
	if t.Right != nil {
		t.Right.Display(depth + 1)
	}
	for i := 0; i < depth; i++ {
		fmt.Print("\t")
	}
	fmt.Println(t.Val)
	if t.Left != nil {
		t.Left.Display(depth + 1)
	}
}

func (t *TreeNode) Delete(val int) error {
	return nil
}

func (t *TreeNode) Length() int {
	return 0
}

func (t *TreeNode) IsPresent(val int) bool {
	return false
}

func (t *TreeNode) IsEmpty() bool {
	return true
}

func (t *TreeNode) IsBST() bool {
	return false
}

func init() {

}
