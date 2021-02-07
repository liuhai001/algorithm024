package main

import (
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

//迭代解法
func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		tmp := make([]int, 0)
		sz := len(stack) //一层的个数
		for sz > 0 {     //输出这一层，并添加子节点
			node := stack[0]
			stack = stack[1:]
			tmp = append(tmp, node.Val)
			for _, child := range node.Children {
				stack = append(stack, child)
			}
			sz--
		}
		res = append(res, tmp)
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
	fmt.Println(levelOrder(root))
}
