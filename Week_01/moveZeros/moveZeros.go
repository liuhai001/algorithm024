package main

import "fmt"

//解法速记：遍历+交换
func moveZeroes(nums []int) {
	j := 0 //j变量记录非0的位置
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 { //i非0，和j位置交换
			nums[j] = nums[i]
			nums[i] = 0
			j++ //j往前挪移一个位置
		}
	}
}
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
