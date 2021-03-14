package main

//贪心算法，从后往前贪心
func canJump(nums []int) bool {
	length := len(nums)
	//初始化可到达位置
	endReach := length - 1
	for i := length - 1; i >= 0; i-- {
		//更新可到达位置
		if nums[i]+i >= endReach {
			endReach = i
		}
	}
	return endReach == 0

}
