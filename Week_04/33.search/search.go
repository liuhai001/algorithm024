package main

//leetcode 33.搜索旋转排序数组
//时间复杂度O(logn) 空间复杂度O(1)
func search(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			//左边有序, 目标在左边
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			//右边有序，目标在右边
			if nums[mid] < target && target <= nums[length-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
