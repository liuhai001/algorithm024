package main

import "sort"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	//为了去重，一定要排序
	sort.Ints(nums)
	n := len(nums)
	digitMap := make([]bool, n)
	var dfs func(int)
	dfs = func(cur int) {
		if cur == n {
			//返回
			res = append(res, append([]int{}, list...))
			return
		}
		for index, val := range nums {
			//关键在这一步, 去重, 同一层有人选过了，不能重复选取
			if digitMap[index] || (index > 0 && val == nums[index-1] && !digitMap[index-1]) {
				continue
			}

			digitMap[index] = true
			list = append(list, val)
			dfs(cur + 1)

			//恢复状态
			digitMap[index] = false
			list = list[:len(list)-1]
		}

	}

	dfs(0)
	return res
}