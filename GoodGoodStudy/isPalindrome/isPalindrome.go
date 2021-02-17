package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//判断是不是回文链表
var left *ListNode

//时间复杂度O(n) 空间复杂度O(n)
func isPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)

	res = left.Val == right.Val && res == true
	left = left.Next
	return res
}

//时间复杂度O(n) 空间复杂度O(n)
func isPalindrome_1(head *ListNode) bool {
	//快慢指针
	slow, fast := head, head
	// slow 指针现在指向链表中点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//如果fast不为nil,说明链表是奇数，slow再往前一步
	if fast != nil {
		slow = slow.Next
	}

	//反转slow为头的链表
	left := head
	right := reverse(slow)
	for right != nil {
		if right.Val != left.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverse(slow *ListNode) *ListNode {
	var pre *ListNode
	cur, nxt := slow, slow
	for cur != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	return pre
}
