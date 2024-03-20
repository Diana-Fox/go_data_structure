package queue

type Queue[T any] interface {
	Enqueue(val T) //入队
	Dequeue() T    //出队
	Empty() bool
}
