/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 */

// @lc code=start
type MyCircularDeque struct {
  data []int
	head int
	tail int
	size int
	cap int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
  mc := MyCircularDeque {
		data: make([]int, k),
		head: 0,
		tail: k - 1,
		size: 0,
		cap: k,
	}
	return mc
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
  if this.IsFull() { return false }
	this.head = (this.head - 1 + this.cap) % this.cap
	this.data[this.head] = value
	this.size++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
  if this.IsFull() { return false }
	this.tail = (this.tail + 1) % this.cap
	this.data[this.tail] = value
	this.size++
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
  if this.IsEmpty() { return false }
	this.size--
	this.head = (this.head + 1) % this.cap
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
  if this.IsEmpty() { return false }
	this.size--
	this.tail = (this.tail - 1 + this.cap) % this.cap
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
  if this.IsEmpty() { return -1 }
	return this.data[this.head]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
  if this.IsEmpty() { return -1 }
	return this.data[this.tail]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
  return this.size == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
  return this.size == this.cap
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
// @lc code=end

