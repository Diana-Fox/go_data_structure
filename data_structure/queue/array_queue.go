package queue

import (
	"github.com/Diana-Fox/go_data_structure/data_structure/list"
)

var q Queue[int] = &ArrayQueue[int]{}

type ArrayQueue[T any] struct {
	date *list.ArrayList[T] //栈完全可以基于list去实现，其实栈只是list限制了一部分功能
}

func (a *ArrayQueue[T]) Size() int {
	return a.date.Size()
}

func (a *ArrayQueue[T]) Front() T {
	value, _ := a.date.Get(0)
	return value
}

func (a *ArrayQueue[T]) End() T {
	value, _ := a.date.Get(a.date.Size() - 1)
	return value
}

func (a *ArrayQueue[T]) IsEmpty() bool {
	return a.date.Size() <= 0
}

// 追加到最后一位
func (a *ArrayQueue[T]) EnQueue(value T) {
	a.date.Append(value)
}

func (a *ArrayQueue[T]) DeQueue() T {
	value, _ := a.date.Delete(a.date.Size() - 1)
	return value
}
