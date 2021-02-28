package main

//返回x 的平方根
//二分查找 1，x 之间对半查找 时间复杂度O(logn)
func mySqrt(x int) int {
	left := 0
	right := x
	res := -1
	for left <= right {
		mid := left + (right-left)>>1
		tmp := mid * mid
		if tmp > x {
			right = mid - 1
		} else if tmp <= x {
			//这里是因为要向下取整
			res = mid
			left = mid + 1
		}
	}
	return res
}
