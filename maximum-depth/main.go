package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func main() {
	var t1 Tree
	t1.insert(4)
	t1.insert(2)
	t1.insert(5)

	fmt.Println(maxDepth(t1.root))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &TreeNode{Val: data}
	} else {
		t.root.insert(data)
	}
}

func (n *TreeNode) insert(data int) {
	if data <= n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: data}
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: data}
		} else {
			n.Right.insert(data)
		}
	}
}
