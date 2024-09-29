package stack

import "github.com/Diana-Fox/go_data_structure/iterator"

// Stack 定义一个栈
type Stack[T any] interface {
	Clear()        //清空栈
	Size() int     //栈大小
	Pop() T        //弹出
	Push(value T)  //押入
	IsEmpty() bool //空栈么
	IsFull() bool  //满栈么
	Iterator() iterator.Iterator[T]
}
