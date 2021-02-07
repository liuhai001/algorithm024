package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//树的前中后序遍历都是递归解法
func inorderTraversal(root *TreeNode) []int {
	//因为返回结果是一个slice，我们要加一层函数，包装一下
	res := make([]int, 0)
	inorder(root, &res)
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}
