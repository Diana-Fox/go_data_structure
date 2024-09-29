package list

import (
	"errors"
)

type ArrayList[T any] struct {
	cap  int
	size int
	data []T //由于go中相关结构为切片实现，切片已经是一个完整数据结构，包含阔缩容等机制
}

//无参,默认大小为10

func NewArrayList[T any]() *ArrayList[T] {
	list := ArrayList[T]{
		size: 0,
		data: make([]T, 10),
	}
	return &list
}

// 只有数组需要
func (a *ArrayList[T]) checkIsFull() {
	if a.size == cap(a.data) { //扩容
		newData := make([]T, a.size, a.size*2) //要先开辟基础内存
		copy(newData, a.data)
		a.data = newData
	}
}
func (a *ArrayList[T]) GetCapSize() int {
	return a.cap
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

func (a *ArrayList[T]) Insert(index int, v T) error {
	if index < 0 || index >= a.size {
		return errors.New("index out of range")
	}
	a.checkIsFull()
	a.data = a.data[:a.size+1] //go的特性，需要先多开辟一个空间
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.size++
	return nil
}

func (a *ArrayList[T]) Set(i int, v T) error {
	if i < 0 || i >= a.size {
		return errors.New("index out of range")
	}
	a.data[i] = v
	return nil
}

func (a *ArrayList[T]) Append(v T) {
	a.data = append(a.data, v)
	a.size += 1 //数据量增加
}

func (a *ArrayList[T]) Clear() {
	a.size = 0
	a.data = make([]T, 0)
}

func (a *ArrayList[T]) Delete(i int) (T, error) { //删除
	var v T
	if a.size < 0 {
		return v, errors.New("index out of range")
	}
	a.size--
	value := a.data[i]
	a.data = append(a.data[:i], a.data[i+1:]...)
	return value, nil
}

func (a *ArrayList[T]) String() string {
	str := "{"
	//for _, v := range a.data {
	//	str = str + fmt.Sprintf("%s,", strconv.Itoa(v)) //默认把所有值都按字符串输出
	//}
	str = str + "}"
	return str
}

// 迭代器
func (a *ArrayList[T]) Iterator() Iterator[T] {
	it := new(ArrayListIterator[T])
	return it
}
