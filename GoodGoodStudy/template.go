package GoodGoodStudy

const MAX_LEVEL = 100

//递归模版
func recursion(level, param int) {
	// recursion terminator 递归终止条件
	if level > MAX_LEVEL {
		// process result
		return
	}

	// process current logic 处理当前逻辑
	process(level, param)

	// drill down 进入下一层
	recursion(level+1, param)

	// reverse the current level status if needed 恢复当前层的状态
}

func process(level, param int) {

}

//DFS模版 递归写法

type Node struct {
	Val      int
	Children []*Node
}

var visited map[int]bool = make(map[int]bool)

func dfs(node *Node) {
	// terminator
	if node == nil {
		return
	}
	if visited[node.Val] {
		// already visited
		return
	}
	visited[node.Val] = true

	// process current node here.
	// ...
	for i := 0; i < len(node.Children); i++ {
		dfs(node.Children[i])
	}
	return
}

//非递归写法：
func dfs_1(node *Node) {
	visited := make(map[int]bool)
	if node == nil {
		return
	}
	stack := []*Node{node}
	for len(stack) > 0 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[tmp.Val] {
			continue
		}
		visited[tmp.Val] = true

		for i := len(tmp.Children) - 1; i >= 0; i-- {
			stack = append(stack, tmp.Children[i])
		}
	}
	return
}

//BFS模版
func bfs(node *Node) {
	visited := make(map[int]bool)
	if node == nil {
		return
	}
	queue := []*Node{node}
	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		if visited[tmp.Val] {
			continue
		}
		visited[tmp.Val] = true

		for i := 0; i <= len(tmp.Children)-1; i++ {
			queue = append(queue, tmp.Children[i])
		}
	}
	return
}
