package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type linkList struct {
	head *ListNode
}

func main() {
	list1 := linkList{}
	list2 := linkList{}

	node11 := &ListNode{Val: 1}
	node12 := &ListNode{Val: 2}
	node13 := &ListNode{Val: 10}
	node21 := &ListNode{Val: 6}
	node22 := &ListNode{Val: 50}
	node23 := &ListNode{Val: 60}
	list1.append(node11)
	list1.append(node12)
	list1.append(node13)
	list2.append(node21)
	list2.append(node22)
	list2.append(node23)
	res := mergeTwoLists2(list1.head, list2.head)
	fmt.Println(res)

}

func (ll *linkList) append(node *ListNode) {

	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}

}

// 2,4  +  3 ------> 2,3,4 RECURSIVE SOLUTION
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list2.Val > list1.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return result.Next
}
