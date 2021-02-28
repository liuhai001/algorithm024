package main

//二分查找变形，时间复杂度O(logn) 空间复杂度O(1)
//注意特殊情况：会数组访问越界，panic！
func findMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	if nums[right] >= nums[left]{
		return nums[0]
	}


	for left <= right {
		mid := left + (right-left)>>1
		//这个在前
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		//这个在后
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		//左边有序，在右边
		if nums[mid] >= nums[0] {
			left = mid + 1
		} else {
			//右边有序，在左边
			right = mid - 1
		}

	}

	return -1
}
