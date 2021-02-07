package main

import "fmt"

func largestRectangleArea(heights []int) int {
	//1.暴力解法 每次找左右边界, 毙掉，超出时间限制，要优化
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		left, right := i, i
		//找左边界
		for j := i - 1; j >= 0; j-- {
			if heights[j] < heights[i] {
				break
			}
			left = j
		}
		//找右边界
		for j := i + 1; j < len(heights); j++ {
			if heights[j] < heights[i] {
				break
			}
			right = j
		}
		maxArea = max((right-left+1)*heights[i], maxArea)

	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea_1(heights []int) int {
	//1.暴力解法 每次找左右边界, 毙掉，超出时间限制，要优化
	maxArea := 0
	queue := make([]int,0)
	for i := 0; i < len(heights); i++ {
		//右边界已经找到
		for len(queue) > 1 && heights[i] < heights[queue[len(queue)-1]] {
			//出栈，计算面积
			height := heights[queue[len(queue)-1]]
			queue = queue[:len(queue)-1]
			right := i
			left := queue[len(queue)-1]
			fmt.Println(left, right, height)
			maxArea = max((right-left-1)*height, maxArea)
		}
		queue = append(queue, i)
	}

	return maxArea
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea_1(heights))
}
