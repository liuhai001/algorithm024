package main

/**
* Definition for a binary tree node.
}
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)

}

func build(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	//寻找根节点
	rootVal := preorder[preStart]

	//寻找根节点在inorder中的位置
	index := inStart
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
		}
	}

	leftSize := index - inStart

	root := &TreeNode{
		Val: rootVal,
	}
	root.Left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1)
	root.Right = build(preorder, preStart+leftSize+1, preEnd, inorder, index+1, inEnd)
	return root
}
