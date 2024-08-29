package list

import (
	"errors"
	"fmt"
)

type ArrayList[T any] struct {
	size int
	data []T //由于go中相关结构为切片实现，切片已经是一个完整数据结构，包含阔缩容等机制
}

//无参,默认大小为

func NewArrayList[T any]() List[T] {
	list := ArrayList[T]{
		size: 0,
		data: make([]T, 0),
	}
	return &list
}

func (a *ArrayList[T]) Size() int {
	return a.size
}

func (a *ArrayList[T]) Get(i int) (T, error) {
	if i < 0 || i >= a.size {
		errors.New("index out of range")
	}
	return a.data[i], nil
}

func (a *ArrayList[T]) Insert(i int, v T) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Set(i int, v T) error {
	if i < 0 || i >= a.size {
		return errors.New("index out of range")
	}
	a.data[i] = v
	return nil
}

func (a *ArrayList[T]) Index(i int, v T) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Append(v T) {
	a.data = append(a.data, v)
	a.size += 1 //数据量增加
}

func (a *ArrayList[T]) Clear() {
	a.size = 0
	a.data = make([]T, 0)
}

func (a *ArrayList[T]) Delete(i int) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) String() string {
	//return fmt.Sprintf(a.data)
	str := "{"
	for _, v := range a.data {
		str = str + fmt.Sprintf("%s,", v) //默认把所有值都按字符串输出
	}
	str = str + "}"
	return str
}
