package list

import (
	"errors"
	"github.com/Diana-Fox/go_data_structure/data_structure/iterator"
)

//var l List[int] = &LoopArray[int]{}

// 循环数组
type LoopArray[T any] struct {
	size  int //长度
	start int //首
	end   int //尾
	data  []T //由于go中相关结构为切片实现，切片已经是一个完整数据结构，包含阔缩容等机制
}

// 循环数组，大小可控
func NewLoopArray[T any](size int) *LoopArray[T] {
	return &LoopArray[T]{
		size:  0,
		start: 0,
		end:   0,
		data:  make([]T, size),
	}
}

func (l *LoopArray[T]) Size() int {
	return l.size
}

func (l *LoopArray[T]) Get(index int) (value T, e error) {
	if index < 0 || index >= l.size { //超过拥有的元素数
		e = errors.New("index out of range")
		return
	}
	//从开始位置偏移若干个位置
	return l.data[l.start+index], e
}

func (l *LoopArray[T]) Set(index int, v T) error {
	l.data[l.start+index] = v
	return nil
}

func (l *LoopArray[T]) Insert(index int, v T) error {
	if index < 0 || index >= l.size { //超过能插入的元素位置了
		return errors.New("index out of range")
	}
	if l.isFull() {
		return errors.New("loop array is full")
	}
	l.end++ //要多一位
	for i := l.end; i > l.start+index; i-- {
		l.data[i] = l.data[i-1]
	}
	l.data[index] = v
	l.size++
	return nil
}

func (l *LoopArray[T]) Append(v T) {
	if l.isFull() {
		return
	}
}
func (l *LoopArray[T]) isFull() bool {
	return l.size > 0 && l.start == l.end-1 //开头追上结尾了
}
func (l *LoopArray[T]) Clear() {
	l.size = 0
	l.start = 0
	l.end = 0
}

func (l *LoopArray[T]) Delete(index int) (T, error) {
	value := l.data[index]
	for i := l.start + index; i < l.end-1; i++ {
		l.data[i] = l.data[i-1]
	}
	l.end--  //结束位置前移
	l.size-- //总个数前移
	return value, nil
}

func (l *LoopArray[T]) String() string {
	//TODO implement me
	panic("implement me")
}

func (l *LoopArray[T]) Iterator() iterator.Iterator[T] {
	it := LoopArrayIterator[T]{
		currentIndex: l.start, //从这个位置开始
	}
	return &it
}
