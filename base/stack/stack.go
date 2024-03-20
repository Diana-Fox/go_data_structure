package stack

import "github.com/Diana-Fox/go_data_structure/base/vector"

type stack[T any] struct {
	//栈是基于vector实现的，但是由于go不支持接口继承接口
	vector vector.Vector[T]
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{}
}

//栈混洗

func (s *stack[T]) Push(val T) {
	s.vector.Put(val) //放到末尾
}

func (s *stack[T]) Top() {
	s.vector.Get(s.vector.Size() - 1) //返回最末尾元素
}

func (s *stack[T]) Pop() T {
	//return s.vector.Remove(s.vector.Size()-1)//vector没实现好，假设有了
	return *new(T)
}

func (s *stack[T]) Size() int {
	return s.vector.Size()
}

func (s *stack[T]) Empty() bool {
	return s.vector.Size() == 0
}

//一般做栈混洗、逆波兰表达式啥的
