package stack

import "github.com/Diana-Fox/go_data_structure/list"

// ArrayStack 数组类型的栈
type ArrayStack[T any] struct {
	date *list.ArrayList[T] //栈完全可以基于list去实现，其实栈只是list限制了一部分功能
}

func NewArrayStack[T any]() *ArrayStack[T] {
	stack := ArrayStack[T]{
		date: list.NewArrayList[T](), //通过arrayList去实现，因为这是个arrayStack
	}
	return &stack
}

func (a *ArrayStack[T]) Clear() {
	a.date.Clear()
}

func (a *ArrayStack[T]) Size() int {
	return a.date.Size()
}

// 每次都删除开头第一个
func (a *ArrayStack[T]) Pop() T {
	value, _ := a.date.Delete(0) //删除第一个
	return value
}

// 每次都追加到最后一个
func (a *ArrayStack[T]) Push(value T) {
	a.date.Append(value)
}

func (a *ArrayStack[T]) IsEmpty() bool {
	return a.date.Size() == 0 //没值
}

// go的切片的一些特性，导致不是很完美
func (a *ArrayStack[T]) IsFull() bool {
	return a.date.Size() == a.date.GetCapSize()
}
