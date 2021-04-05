package main

import (
	"fmt"
)

//归并排序, 稳定排序，最好O(nlogn),最坏O(nlogn),平均O(nlogn)
//空间复杂度，O(n)
func MergeSort(nums []int) {
	mergeSortC(nums, 0, len(nums)-1)
}

func mergeSortC(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)>>1
	mergeSortC(nums, start, mid)
	mergeSortC(nums, mid+1, end)
	merge(nums, start, mid, end)
}
func merge(nums []int, start, mid, end int) {
	//申请临时空间
	tmp := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			k++
			i++
		} else {
			tmp[k] = nums[j]
			k++
			j++
		}
	}

	//将剩余的数据拷贝到tmp
	p := i
	q := mid
	if j <= end {
		p = j
		q = end
	}

	for p <= q {
		tmp[k] = nums[p]
		k++
		p++
	}

	//赋值回去
	for i, v := range tmp {
		nums[start+i] = v
	}

}

//快速排序，不稳定排序，最好O(nlogn),最坏O(n^2),平均O(nlogn)
//原地排序，空间复杂度O(1)
func QuickSort(nums []int) {
	quickSortC(nums, 0, len(nums)-1)
}

func quickSortC(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(nums, start, end) //获取分区点
	quickSortC(nums, start, pivot-1)
	quickSortC(nums, pivot+1, end)
}

func partition(nums []int, start, end int) int {
	//选择最后一个元素，当分区值，start~ j-1 小于pivot，j+1 到end大于pivot
	pivot := nums[end]
	j := start
	for i := start; i <= end-1; i++ {
		if nums[i] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	//分区值放中间
	nums[j], nums[end] = nums[end], nums[j]
	return j
}

func main() {
	ans := []int{1, 1, 6, 3, 0, 5, 2, 3, 5, 6, 11, 2, 3, 55}
	QuickSort(ans)
	fmt.Println(ans)
}
