package main

import "math/bits"

var solutions [][]string

//方法一：用三个map存 列撇捺 是否占用
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

//方法二、 用是那个整数来存列撇捺 是否占用
func solveNQueens_1(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	solve(queens, n, 0, 0, 0, 0)
	return solutions
}

func solve(queens []int, n, row, columns, diagonals1, diagonals2 int) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	//三个整数或后取反，bit 为1 的就是可以放皇后的位置
	availablePositions := ((1 << n) - 1) & (^(columns | diagonals1 | diagonals2))
	for availablePositions != 0 {
		//拿出最低位可以放的位置
		position := availablePositions & (-availablePositions)
		//遍历过的位，给它置0
		availablePositions = availablePositions & (availablePositions - 1)
		//计算所在列
		column := bits.OnesCount(uint(position - 1))
		queens[row] = column
		solve(queens, n, row+1, columns|position, (diagonals1|position)>>1, (diagonals2|position)<<1)
		queens[row] = -1
	}
}
