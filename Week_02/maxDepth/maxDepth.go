package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//解题思路1：递归思想，左子树，右子树都是一个完整树，可以递归
func maxDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth_1(root.Left), maxDepth_1(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//解题思路2：树的最大深度就是有多少层，可以用层序遍历，用队列实现
func maxDepth_2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root) //先把根节点放入队列
	for len(queue) > 0 {        //判断队列里面是否有节点
		sz := len(queue)
		for sz > 0 { //一层全部出列
			//根节点出列
			root := queue[0]
			queue = queue[1:]
			//把左右子节点加入列
			if root.Left != nil {
				queue = append(queue, root.Left)
			}

			if root.Right != nil {
				queue = append(queue, root.Right)
			}
			sz--
		}

		//遍历完一层+1
		depth++
	}

	return depth
}
