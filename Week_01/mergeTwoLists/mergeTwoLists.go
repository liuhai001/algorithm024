package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists_1(l1 *ListNode, l2 *ListNode) *ListNode {
	//1.迭代解法
	dummyHead := &ListNode{}
	prev := &ListNode{}
	dummyHead = prev
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return dummyHead.Next
}

func mergeTwoLists_2(l1 *ListNode, l2 *ListNode) *ListNode {
	//1.递归解法
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists_2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists_2(l1, l2.Next)
		return l2
	}

	return nil
}

func buildLinkList(nums []int) *ListNode {
	dummyHead := &ListNode{}
	head := dummyHead
	for _, value := range nums {
		newNode := &ListNode{
			Val: value,
		}
		head.Next = newNode
		head = head.Next

	}
	return dummyHead.Next
}

func (head *ListNode) printLinkList() []int {
	nums := make([]int, 0)
	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}

	return nums
}

func main() {
	l1 := buildLinkList([]int{1, 2, 4})
	l2 := buildLinkList([]int{1, 3, 4})
	fmt.Println(mergeTwoLists_1(l1, l2).printLinkList())

	l3 := buildLinkList([]int{1, 5, 7})
	l4 := buildLinkList([]int{2, 4, 8})
	fmt.Println(mergeTwoLists_2(l3, l4).printLinkList())
}
