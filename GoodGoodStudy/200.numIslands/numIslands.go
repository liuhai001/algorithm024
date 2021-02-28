package main

//DFS 思想：遇到1，count +1, 并把所有上下左右1夷为平地
//行，列 深度搜索
//时间复杂度O(MN) 空间复杂度O(MN)
func numIslands(grid [][]byte) int {
	//行 列
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	count := 0
	var dfs func([][]byte, int, int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '0'
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}
