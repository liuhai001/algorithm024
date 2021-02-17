package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	a := head
	b := head

	//递归base case 不满足k个节点原数据返回
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

//反转a到b节点的链表[a,b)
func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	nxt := a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
