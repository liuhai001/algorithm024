package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

//递归解法
func preorder(root *Node) []int {
	res := make([]int, 0)
	preorderNtrees(root, &res)
	return res
}

func preorderNtrees(root *Node, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	for _, child := range root.Children {
		preorderNtrees(child, res)
	}
}

//迭代解法
func preorder_2(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return res
}

func main() {
	root := &Node{Val: 1}
	node3 := &Node{Val: 3}
	node2 := &Node{Val: 2}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	root.Children = []*Node{node3, node2, node4}
	node3.Children = []*Node{node5, node6}
	fmt.Println(preorder_2(root))
}
