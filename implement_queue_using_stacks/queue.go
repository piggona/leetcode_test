package main

import "fmt"

type Stack []int

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *Stack) Heap() int {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

type MyQueue struct {
	stackIn  *Stack
	stackOut *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  NewStack(),
		stackOut: NewStack(),
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	if !this.stackOut.IsEmpty() {
		return this.stackOut.Pop()
	}
	for !this.stackIn.IsEmpty() {
		this.stackOut.Push(this.stackIn.Pop())
	}
	return this.stackOut.Pop()
}

func (this *MyQueue) Peek() int {
	if !this.stackOut.IsEmpty() {
		return this.stackOut.Heap()
	}
	for !this.stackIn.IsEmpty() {
		this.stackOut.Push(this.stackIn.Pop())
	}
	return this.stackOut.Heap()
}

func (this *MyQueue) Empty() bool {
	return this.stackIn.IsEmpty() && this.stackOut.IsEmpty()
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)

	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
}
