package main

import "container/heap"

//方法1: 建立一个k个元数的小顶堆 时间复杂度O(n) + O(nlogk) 空间复杂度 O(n) + O(k)
func topKFrequent(nums []int, k int) []int {
	//统计频次 时间复杂度O(n), 空间O(n)
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	//堆空间O(k) O(nlogk)
	//建立堆，因为是数和频率两个指标，所以是一个两元素数组的slice
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

//方法2: 建立一个k个元数的小顶堆
func topKFrequent(nums []int, k int) []int {
	//统计频次
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}

	//建立堆，因为是数和频率两个指标，所以是一个两元素数组的slice
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		if h.Len() == k {
			if h.GetTop().([2]int)[1] < value {
				heap.Pop(h)
				heap.Push(h, [2]int{key, value})
			}
		} else {
			heap.Push(h, [2]int{key, value})
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] } //i是子节点，j是父亲节点，Less返回true交换
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IHeap) GetTop() interface{} {
	if h.Len() > 0 {
		old := *h
		return old[0]
	}
	return []int{}
}
