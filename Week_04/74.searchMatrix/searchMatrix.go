package main

//把每一行接起来，拼成一个一维数组，然后二分查找
func searchMatrix(matrix [][]int, target int) bool {
	//二维数组可以虚拟成一维数组
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	left := 0
	right := m*n - 1
	for left <= right {
		mid := left + (right-left)>>1
		tmp := matrix[mid/n][mid%n]
		if tmp == target {
			return true
		} else if tmp > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return false
}
