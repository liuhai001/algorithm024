package main

import "fmt"

//解法速记：hash + 遍历
func twoSum(nums []int, target int) []int {
	//数组的遍历，每次都找target - nums[i]的数是否存在
	//使用一个map
	numMap := make(map[int]int, 0)
	for i, num := range nums {
		if index, ok := numMap[target-num]; ok {
			return []int{i, index}
		}
		numMap[num] = i
	}

	return []int{}
}
func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
