package list

type ListNode[T any] interface {
	Pred() ListNode[T]          //前驱
	Succ() ListNode[T]          //后继
	data() T                    //节点数据
	InsertAsPred(p ListNode[T]) //插入前驱节点
	InsertAsSucc(p ListNode[T]) //插入后继节点
}
type Node[T any] struct {
	pred *Node[T] //前驱节点
	succ *Node[T] //后继节点
	val  T
}

func NewNode[T any](e T) ListNode[T] {
	return &Node[T]{
		val: e,
	}
}
func (n *Node[T]) Pred() ListNode[T] {
	return n.pred
}

func (n *Node[T]) Succ() ListNode[T] {
	return n.succ
}

func (n *Node[T]) data() T {
	return n.val
}

func (n *Node[T]) InsertAsPred(p ListNode[T]) {
	//nn := p.(*Node[T])

}

func (n *Node[T]) InsertAsSucc(p ListNode[T]) {
	//nn := p.(*Node[T])

}
