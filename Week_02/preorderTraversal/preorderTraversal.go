package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//树的前中后序遍历都是递归解法
func preorderTraversal(root *TreeNode) []int {
	//因为返回结果是一个slice，我们要加一层函数，包装一下
	res := make([]int, 0)
	preorder(root, &res)
	return res
}

func preorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorder(root.Left, res)
	preorder(root.Right, res)
}

