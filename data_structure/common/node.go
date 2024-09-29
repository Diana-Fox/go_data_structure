package common

// 标准node节点
type Node[T any] struct {
	value string
	prev  *Node[T] //上一个
	next  *Node[T] //下一个
}
