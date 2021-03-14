package main

import "sort"

//贪心算法，排好序之后，给每个孩子找最小满足它的饼干
//时间复杂度O(nlogn+mlogm)
func findContentChildren(g []int, s []int) int {
	//排序
	sort.Ints(g)
	sort.Ints(s)
	numOfGreed := len(g)
	numOfS := len(s)
	count := 0
	for i, j := 0, 0; i < numOfGreed && j < numOfS; i++ {
		//找最小满足孩子的饼干
		for j < numOfS && g[i] > s[j] {
			j++
		}
		if j < numOfS {
			count++
			j++
		}
	}

	return count
}
