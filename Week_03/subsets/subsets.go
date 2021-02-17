package main

//分治解法
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	recursion(0, len(nums), nums, &res, []int{})
	return res
}

func recursion(level, max int, nums []int, res *[][]int, list []int) {
	if level == max {
		//append注意，用新生成的list
		*res = append(*res, append([]int{}, list...))
		return
	}

	//不包括当前元素，往下一层
	recursion(level+1, max, nums, res, list)

	//包括当前元素往下一层
	list = append(list, nums[level])
	recursion(level+1, max, nums, res, list)

	//恢复当前状态
	list = list[:len(list)-1]
}
