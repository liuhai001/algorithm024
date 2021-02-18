package main

//回溯算法思想：三个位置，每一个位置可以选1，2，3
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	digitMap := make(map[int]bool)
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			//返回
			res = append(res, append([]int{}, list...))
			return
		}
		for _, val := range nums {
			if digitMap[val] {
				continue
			}

			digitMap[val] = true
			list = append(list, val)
			dfs(cur + 1)

			//恢复状态
			delete(digitMap,val)
			list = list[:len(list)-1]
		}

	}

	dfs(0)
	return res
}
