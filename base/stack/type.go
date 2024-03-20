package stack

// 栈
type Stack[T any] interface {
	Push(val T)  //推入
	Top()        //查询顶部
	Pop() T      //弹出
	Size() int   //栈数量
	Empty() bool //为空
}
