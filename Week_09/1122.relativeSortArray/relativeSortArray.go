package main

import "sort"

//自定义排序规则
func relativeSortArray(arr1 []int, arr2 []int) []int {
	rankMap := make(map[int]int)
	for i, v := range arr2 {
		rankMap[v] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, hasX := rankMap[x]
		rankY, hasY := rankMap[y]
		if hasX && hasY {
			return rankX < rankY
		}
		if hasX || hasY {
			return hasX
		}
		return x < y
	})

	return arr1
}
