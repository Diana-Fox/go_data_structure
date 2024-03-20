package queue

type queue[T any] struct {
}

func (q *queue[T]) Enqueue(val T) {
	//TODO implement me
	panic("implement me")
}

func (q *queue[T]) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Dequeue() T {
	//TODO implement me
	panic("implement me")
}
