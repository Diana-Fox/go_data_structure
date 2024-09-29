package queue

// 队列
type Queue[T any] interface {
	Size() int
	Front() T        //第一个元素
	End() T          //最后一个元素
	IsEmpty() bool   //是否为空
	EnQueue(value T) //入队
	DeQueue() T      //出队
}
