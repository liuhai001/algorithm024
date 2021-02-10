package main

import (
	"fmt"
)

//二叉堆的实现，以及堆排序
type Heap struct {
	a     []int //底层一维数组
	count int   //堆中元素个数
}

//数组从0开始，左孩子右孩子 2*i+1 2*i+2  父亲节点 （j-1）/2
//非叶子节点 （j-2）/2

//初始化建堆
func Init(nums []int) *Heap {
	h := &Heap{
		a:     nums,
		count: len(nums),
	}

	//建堆，时间复杂度O(n)
	h.buildHeap()
	return h
}

func (h *Heap) buildHeap() {
	//因为是从上到下，叶子节点不需要处理，只需要处理非叶子节点，一半节点
	for i := (h.Len()-1)/2 - 1; i >= 0; i-- {
		h.down(i, h.Len())
	}
}

//从上到下
func (h *Heap) down(i0 int, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1 //左孩子
		if j1 >= n || j1 < 0 {
			break
		}
		pos := j1
		if j2 := j1 + 1; j2 < n && h.Less(j1, j2) {
			pos = j2 //右孩子
		}
		if h.Less(pos, i) {
			break
		}

		h.swap(i, pos)
		i = pos
	}
	return i > i0
}

//从下到上
func (h *Heap) up(i0 int) bool {
	i := i0
	for {
		j1 := (i - 1) / 2 //父亲节点
		if j1 == i || !h.Less(j1, i) {
			break
		}

		h.swap(i, j1)
		i = j1
	}
	return i < i0
}

func (h *Heap) Len() int {
	return h.count
}

func (h *Heap) Less(i, j int) bool {
	return h.a[i] < h.a[j] //小于就是大顶堆，大于就是小顶堆
}

func (h *Heap) swap(i, j int) {
	h.a[i], h.a[j] = h.a[j], h.a[i]
}

func Sort(a []int) {
	h := Init(a)
	k := h.Len()
	for k > 1 {
		h.swap(0, k-1)
		k--
		h.down(0, k)
	}
}

func main() {
	a := []int{2, 1, 5, 3, 8}
	Sort(a)
	fmt.Println(a)
}
