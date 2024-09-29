package list

import "github.com/Diana-Fox/go_data_structure/data_structure/iterator"

var l iterator.Iterator[int] = &LoopArrayIterator[int]{}

type LoopArrayIterator[T any] struct {
	list         *LoopArray[T]
	currentIndex int
}

func (l *LoopArrayIterator[T]) HasNext() bool {
	return l.currentIndex == l.list.end
}

func (l *LoopArrayIterator[T]) Next() (T, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LoopArrayIterator[T]) Remove() error {
	//TODO implement me
	panic("implement me")
}

func (l *LoopArrayIterator[T]) GetIndex() int {
	//TODO implement me
	panic("implement me")
}
