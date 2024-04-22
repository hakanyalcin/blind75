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
	var t1, t2 Tree
	t1.insert(4)
	t1.insert(2)
	t1.insert(5)

	t2.insert(4)
	t2.insert(2)
	t2.insert(5)
	fmt.Println(isSameTree(t1.root, t2.root))

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	return left && right
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
