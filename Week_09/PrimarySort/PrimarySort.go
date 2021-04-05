package main

import "fmt"

//选择排序，不稳定排序，最好O(n^2),最坏O(n^2),平均O(n^2)
func SelectionSort(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		minIndex := i //默认最小元素位置为i
		//找出最小元素
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		//交换
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

//插入排序, 稳定排序，最好O(n),最坏O(n^2),平均O(n^2)
func InsertionSort(nums []int) {
	length := len(nums)
	for i := 1; i < length; i++ {
		//数组i前面已经有序了
		j := i - 1
		current := nums[i]
		//数据往后移动
		for j >= 0 && nums[j] > current {
			nums[j+1] = nums[j]
			j--
		}
		//找到了坑位，插入
		nums[j+1] = current
	}
}

//冒泡排序, 稳定排序，最好O(n),最坏O(n^2),平均O(n^2)
func BubbleSort(nums []int) {
	length := len(nums)
	//只用遍历length-1躺
	for i := 0; i < length-1; i++ {
		//提前退出冒泡循环的标志位
		flag := false
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] { //相邻元素两两交换
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag { //没有数据交换，提前退出
			break
		}
	}
}

func main() {
	ans := []int{1, 8, 6, 3, 3, 5, 2, 3, 5, 6, 11, 2, 3, 55}
	InsertionSort(ans)
	fmt.Println(ans)
}
