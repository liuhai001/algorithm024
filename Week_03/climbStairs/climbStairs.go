package main

//递归解法
func climbStairs(n int) int {
	statusMap := make(map[int]int, 0)
	return process(statusMap, n)
}

func process(m map[int]int, n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	if val, ok := m[n]; ok {
		return val
	}

	value := process(m, n-1) + process(m, n-2)
	m[n] = value
	return value
}

//递推解法
func climbStairs_1(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	ppre, pre := 1, 2
	sum := 0

	for i := 3; i <= n; i++ {
		sum = ppre + pre
		ppre = pre
		pre = sum
	}

	return sum
}
