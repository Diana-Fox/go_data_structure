package list

import (
	"errors"
)

type ArrayListIterator[T any] struct {
	list         List[T] //类的包含
	currentIndex int     //当前索引
}

// 迭代器
func NewArrayListIterator[T any](list List[T]) Iterator[T] {
	return &ArrayListIterator[T]{
		list:         list,
		currentIndex: 0,
	}
}

func (a *ArrayListIterator[T]) HashNext() bool {
	return a.currentIndex < a.list.Size() //当前迭代是否小于list元素个数
}

func (a *ArrayListIterator[T]) Next() (T, error) {
	if !a.HashNext() {
		return *new(T), errors.New("没有下一个元素了")
	}
	e, err := a.list.Get(a.currentIndex)
	a.currentIndex++
	return e, err
}

func (a *ArrayListIterator[T]) Remove() {
	a.currentIndex--
	a.list.Delete(a.currentIndex)
}
