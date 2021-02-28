package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


//BFS解法 时间复杂度O(n) 空间复杂度O(n)
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		//记录这一层的个数，等下统一出队
		num := len(queue)
		tmpRes := make([]int, 0)
		for num > 0 {
			tmp := queue[0]
			queue = queue[1:]
			tmpRes = append(tmpRes, tmp.Val)

			if tmp.Left != nil {
				queue = append(queue, tmp.Left)

			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)

			}
			num--
		}
		res = append(res, tmpRes)
	}
	return res
}
