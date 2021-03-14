package main

//贪心算法
//每次都更新最大欧式距离，只要没有碰到障碍物就更新
func robotSim(commands []int, obstacles [][]int) int {
	//上左下右 四个方向
	dx := []int{0, -1, 0, 1}
	dy := []int{1, 0, -1, 0}
	x, y, di := 0, 0, 0 //初始化原点（0，0） 方向上
	res := 0
	//障碍存到map中，方便快速确认
	obMap := make(map[[2]int]bool, 0)
	for _, tmp := range obstacles {
		obMap[[2]int{tmp[0], tmp[1]}] = true
	}

	for _, cmd := range commands {
		if cmd == -1 {
			di = (di + 3) % 4
		} else if cmd == -2 {
			di = (di + 1) % 4
		} else {
			for i := 0; i < cmd; i++ {
				//往di 方向走一步
				px := x + dx[di]
				py := y + dy[di]
				//不是障碍点
				if !obMap[[2]int{px, py}] {
					x = px
					y = py
				} else {
					//是障碍点，结束这次命令
					break
				}

			}
		}
		res = max(res, x*x+y*y)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
