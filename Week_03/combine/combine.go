package main

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var dfs func(int)
	list := make([]int, 0)
	dfs = func(cur int) {
		// 剪枝：list 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(list)+(n-cur+1) < k {
			return
		}

		if len(list) == k {
			res = append(res, append([]int{}, list...))
			return
		}

		if cur == n+1 {
			return
		}

		//加这个元素
		list = append(list, cur)
		dfs(cur + 1)

		//不加这个元素
		list = list[:len(list)-1]
		dfs(cur + 1)
	}

	dfs(1)
	return res
}
