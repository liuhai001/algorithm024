package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack    []int //栈
	minStack []int //最小栈
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}

	stack.minStack = append(stack.minStack, math.MaxInt32)
	return stack
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.minStack = append(this.minStack, min(x, this.minStack[len(this.minStack)-1]))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	return  this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(-1)
	obj.Push(4)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_3, param_4)
}
