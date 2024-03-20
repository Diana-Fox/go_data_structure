package tree

// 树节点
type Node[T any] interface {
	InsertAsLc(val T) //作为左孩子插入
	InsertAsRc(val T) //作为右孩子插入
	Succ()            //（中序意义下）当前节点的直接后继
}
type BinNode[T any] struct {
	lChild *BinNode[T] //左孩子
	parent *BinNode[T] //父节点
	rChild *BinNode[T] //右孩子
	data   T           //数据
	height int         //深度
	size   int         //子树规模
}

func NewBinNode[T any](data T, parent *BinNode[T]) *BinNode[T] {
	return &BinNode[T]{
		data:   data,
		parent: parent,
	}
}

func (b *BinNode[T]) InsertAsLc(val T) {
	nn := NewBinNode(val, b) //新节点
	b.lChild = nn
}

func (b *BinNode[T]) InsertAsRc(val T) {
	nn := NewBinNode(val, b) //新节点
	b.rChild = nn
}

// 统计后代个数,会遍历所有节点
func (b *BinNode[T]) Size() int {
	s := 1
	lChild := b.lChild
	for lChild != nil {
		s += lChild.Size() //左边规模
	}
	rChild := b.rChild
	for rChild != nil {
		s += lChild.Size() //右边规模
	}
	return s
}
func (b *BinNode[T]) Succ() {
	//TODO implement me
	panic("implement me")
}
