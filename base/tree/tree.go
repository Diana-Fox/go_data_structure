package tree

import (
	"github.com/Diana-Fox/go_data_structure/base/queue"
	"github.com/Diana-Fox/go_data_structure/base/stack"
)

type binTree[T any] struct {
	size int         //树的规模
	root *BinNode[T] //根节点
}

// 左边插入
func (b *binTree[T]) InsertAsRC(node *BinNode[T], val T) *BinNode[T] {
	b.size++
	node.InsertAsLc(val)
	b.updateHeightAbove(node)
	return node.rChild
}

// 更新x节点的高度
func (b *binTree[T]) updateHeight(x *BinNode[T]) {
	lSize := b.stature(x.lChild) //左边高度
	rSize := b.stature(x.rChild) //右边高度
	height := 1
	if lSize > rSize {
		height += lSize
	} else {
		height += rSize
	}
	x.height = height
}

// 更新X和祖先的高度
func (b *binTree[T]) updateHeightAbove(x *BinNode[T]) {
	for x.parent != nil {
		b.updateHeightAbove(x)
		x = x.parent
	}
}
func (b *binTree[T]) travLevel(visit func(val T)) {
	q := queue.NewQueue[*BinNode[T]]()
	q.Enqueue(b.root) //根节点入队
	for !q.Empty() {
		x := q.Dequeue()
		visit(x.data)
		if x.lChild != nil {
			q.Enqueue(x.lChild)
		}
		if x.rChild != nil {
			q.Enqueue(x.rChild)
		}
	}
}

// 中序遍历，主要把左孩子放进去，没有了就去访问然后去嚯嚯右子树的左孩子
func (b *binTree[T]) traverseIn(x *BinNode[T], visit func(val T)) {
	stack := stack.NewStack[*BinNode[T]]()
	for true {
		b.goAlongLeftBranch(x, stack)
		if stack.Empty() {
			break
		}
		x = stack.Pop() //x的左子树没有了
		visit(x.data)
		x = x.rChild //再去右子树
	}
}

// 沿着左分支，一直往下
func (b *binTree[T]) goAlongLeftBranch(x *BinNode[T], s stack.Stack[*BinNode[T]]) {
	for x != nil {
		s.Push(x)
		x = x.lChild
	}
}

// 极致的性能,先序遍历
func (b *binTree[T]) traverseC(x *BinNode[T], visit func(val T)) {
	stack := stack.NewStack[*BinNode[T]]()
	for true {
		b.visitAlongLeftBranch(x, visit, stack) //访问左边
		if stack.Empty() {
			break
		}
		x = stack.Pop() //弹出右边
	}
}

// 先序遍历，主要是把右孩子放进去
func (b *binTree[T]) visitAlongLeftBranch(x *BinNode[T], visit func(val T), s stack.Stack[*BinNode[T]]) {
	for x != nil {
		visit(x.data)
		s.Push(x.rChild) //右边入队
		x = x.lChild     //沿着左边继续
	}
}

// 尾递归，所以可以改成迭代,这个方法，左右子树都入栈了
func (b *binTree[T]) traverseB(x *BinNode[T], visit func(val T)) {
	stack := stack.NewStack[*BinNode[T]]()
	if x != nil {
		stack.Push(x)
	}
	for !stack.Empty() {
		n := stack.Pop()
		visit(n.data)
		if n.lChild != nil {
			stack.Push(n.lChild)
		}
		if n.rChild != nil {
			stack.Push(n.rChild)
		}
	}
}

// 传入遍历方法,迭代的方式
func (b *binTree[T]) traverseA(x *BinNode[T], visit func(val T)) {
	if x == nil {
		return
	}
	visit(x.data)
	b.traverseA(x.lChild, visit)
	b.traverseA(x.rChild, visit)
}
func (b *binTree[T]) stature(child *BinNode[T]) int {
	return 1
}
