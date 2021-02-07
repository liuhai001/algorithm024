package main

import "fmt"

type MyCircularDeque struct {
	front int   //头
	rear  int   //尾
	cap   int   //容量
	array []int //底层数组
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		front: 0,
		rear:  0,
		cap:   k + 1,
		array: make([]int, k+1, k+1),
	}

}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.front = (this.front - 1 + this.cap) % this.cap
	this.array[this.front] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.array[this.rear] = value
	this.rear = (this.rear + 1) % this.cap
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.cap
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = (this.rear - 1 + this.cap) % this.cap
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.array[this.front]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.array[(this.rear-1+this.cap)%this.cap]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.front == this.rear {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if (this.rear+1)%this.cap == this.front {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func main() {
	obj := Constructor(4)
	param_1 := obj.InsertFront(1)
	param_2 := obj.InsertLast(3)
	param_3 := obj.DeleteFront()
	param_4 := obj.DeleteLast()
	param_5 := obj.GetFront()
	param_6 := obj.GetRear()
	param_7 := obj.IsEmpty()
	param_8 := obj.IsFull()

	fmt.Println(param_1, param_2, param_3, param_4, param_5, param_6, param_7, param_8)
}
