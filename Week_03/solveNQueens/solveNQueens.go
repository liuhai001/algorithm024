package main

var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}

	//初始化皇后的位置
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	//定义三个set保存已经有皇后的列，撇，捺
	columns := map[int]bool{}
	diagonals1 := map[int]bool{}
	diagonals2 := map[int]bool{}
	backtrack(queens, n, 0, columns, diagonals1, diagonals2)
	return solutions
}

func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	for i := 0; i < n; i++ {

		//数据处理
		if columns[i] {
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		//进入下一层
		backtrack(queens, n, row+1, columns, diagonals1, diagonals2)

		//恢复状态
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	//行
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		//列
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
