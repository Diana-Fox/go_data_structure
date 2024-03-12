package list

import (
	"errors"
	"fmt"
)

type ArrayList[T any] struct {
	dataStore []T //使用泛型
	theSize   int
}

func NewList[T any]() List[T] {
	return &ArrayList[T]{
		dataStore: make([]T, 0, 10),
		theSize:   0,
	}
}
func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		dataStore: make([]T, 0, 10),
		theSize:   0,
	}
}
func (a *ArrayList[T]) Size() int {
	return a.theSize
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || a.theSize < index {
		return *new(T), errors.New("索引越界")
	}
	return a.dataStore[index], nil
}

func (a *ArrayList[T]) Set(index int, value T) error {
	if index < 0 || a.theSize < index {
		return errors.New("索引越界")
	}
	a.dataStore[index] = value
	return nil
}

func (a *ArrayList[T]) Insert(index int, newValue T) error {
	if index < 0 || a.theSize < index {
		return errors.New("索引越界")
	}
	a.checkIsFull() //看是否要扩容
	a.dataStore = a.dataStore[:a.theSize+1]
	for i := a.theSize; i > index; i-- {
		a.dataStore[i] = a.dataStore[i-1]
	}
	a.dataStore[index] = newValue
	return nil
}
func (a *ArrayList[T]) checkIsFull() {
	if a.theSize == cap(a.dataStore) {
		//判断内存的使用，超出部分进行扩容,扩大2倍
		newDataStore := make([]T, a.theSize, a.theSize*2)
		//for i := 0; i < len(a.dataStore); i++ {
		//	newDataStore[i] = a.dataStore[i]
		//}
		copy(newDataStore, a.dataStore)
		a.dataStore = newDataStore
	}
}
func (a *ArrayList[T]) Append(newValue T) {
	a.dataStore = append(a.dataStore, newValue)
	a.theSize++
}

func (a *ArrayList[T]) Clear() {
	a.theSize = 0
	a.dataStore = make([]T, 0, 10)
}

func (a *ArrayList[T]) Delete(index int) error {
	a.dataStore = append(a.dataStore[:index], a.dataStore[index+1:]...)
	a.theSize--
	return nil
}

func (a *ArrayList[T]) String() string {
	return fmt.Sprintln(a.dataStore)
}

// 实现Iterable接口
func (a *ArrayList[T]) Iterator() Iterator[T] {
	return NewArrayListIterator[T](a)
}
