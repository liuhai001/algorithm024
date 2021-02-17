package main

import "fmt"

//分治解法
func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return recursion(x, n)
}

func recursion(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	var res float64
	subRes := recursion(x, n/2)
	if n%2 == 0 {
		//偶数
		res = subRes * subRes
	} else {
		//奇数
		res = subRes * subRes * x
	}

	return res
}

func main() {
	fmt.Println(myPow(2, -2))
}
