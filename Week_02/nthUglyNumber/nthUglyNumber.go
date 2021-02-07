package main

import "fmt"

//丑数, 动态规划
func nthUglyNumber(n int) int {
	dp := make([]int, n, n)
	a, b, c := 0, 0, 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(dp[a]*2, dp[b]*3, dp[c]*5)
		if dp[i] == dp[a]*2 {
			a += 1
		}

		if dp[i] == dp[b]*3 {
			b += 1
		}

		if dp[i] == dp[c]*5 {
			c += 1
		}
	}

	return dp[n-1]
}
func min(a, b, c int) int {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}

	if b > c {
		b, c = c, b
	}

	return a
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
