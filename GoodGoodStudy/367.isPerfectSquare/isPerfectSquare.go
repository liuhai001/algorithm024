package main

//二分查找 1，x 之间二分，相等即是完全平方数, 时间复杂度O(logn)
func isPerfectSquare(num int) bool {
	left := 1
	right := num
	for left <= right {
		mid := left + (right-left)>>1
		tmp := mid * mid
		if tmp > num {
			right = mid - 1
		} else if tmp < num {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}
