package list

import (
	"github.com/Diana-Fox/go_data_structure/common"
	"github.com/Diana-Fox/go_data_structure/iterator"
)

type LinkList[T any] struct {
	size int
	head *common.Node[T] //头节点
	tail *common.Node[T] //尾节点
}

func NewLinkList[T any]() *LinkList[T] {
	list := LinkList[T]{}
	return &list
}

func (l *LinkList[T]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (l *LinkList[T]) Get(index int) (T, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkList[T]) Set(index int, v T) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkList[T]) Insert(index int, v T) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkList[T]) Append(v T) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkList[T]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (l *LinkList[T]) Delete(i int) (T, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkList[T]) String() string {
	//TODO implement me
	panic("implement me")
}

func (l *LinkList[T]) Iterator() iterator.Iterator[T] {
	//TODO implement me
	panic("implement me")
}
