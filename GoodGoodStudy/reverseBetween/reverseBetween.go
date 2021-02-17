package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	start := pre.Next  // a pointer to the beginning of a sub-list that will be reversed
	then := start.Next // a pointer to a node that will be reversed

	// 1 - 2 -3 - 4 - 5 ; m=2; n =4 ---> pre = 1, start = 2, then = 3
	// dummy-> 1 -> 2 -> 3 -> 4 -> 5

	for i := 0; i < n-m; i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}

	// first reversing : dummy->1 - 3 - 2 - 4 - 5; pre = 1, start = 2, then = 4
	// second reversing: dummy->1 - 4 - 3 - 2 - 5; pre = 1, start = 2, then = 5 (finish)
	return dummy.Next
}

func reverseBetween_1(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

//反转链表前n个节点
var successor *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

//反转head到nil的链表
func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	nxt := head
	for cur != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

//反转a到b节点的链表[a,b)
func reverse_1(a, b *ListNode) *ListNode {
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
