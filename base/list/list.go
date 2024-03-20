package list

import "errors"

type list[T any] struct {
	size    int
	header  ListNode[T]        //头节点,是哨兵
	trailer ListNode[T]        //尾节点，哨兵
	equal   func(v1, v2 T) int //比较方法
}

func (l *list[T]) Size() int {
	return l.size
}

func (l *list[T]) Disordered() {
	//TODO implement me
	panic("implement me")
}

func (l *list[T]) Sort() {
	//TODO implement me
	panic("implement me")
}

// 选择排序，主要是在左侧无序的部分找到最大，放入右侧有序部分
func (l *list[T]) selectionSort() {
	size := l.size
	head := l.header
	tail := l.trailer
	for i := 0; i < size; i++ {
		tail = tail.Succ() //
		for 1 < size {
			l.InsertBefore(tail, l.Remove(l.selectMax(head.Succ(), i)))
			tail = tail.Pred()
			size--
		}
	}
}

// 在n个前驱中找出最大的元素，遍历取max即可
func (l *list[T]) selectMax(succ ListNode[T], i int) ListNode[T] {
	return nil
}

// 插入排序，左边为有序区间，右边为无序区间，新元素，找到合适的位置，把大于元素的右移一个位置
// 再把元素插入对应位置，最终完全有序。最好的情况下是O(n)，最坏是O(n^2）
func (l *list[T]) insertSort() {
	root, _ := l.First()
	size := l.size
	for i := 0; i < size; i++ {
		//Search接口应该设计成通用的，但是没按通用的写，所以这里暂时这样了
		//l.InsertAfter(l.Search(root.data(),i,root),root.data())
		root = root.Succ()
		l.Remove(root.Pred())
	}
}
func (l *list[T]) Find(e T) ListNode[T] {
	size := l.size
	p, _ := l.First()
	for size > 0 {
		if l.equal(p.data(), e) == 0 {
			return p
		}
	}
	return nil
}

func (l *list[T]) Search(e T) ListNode[T] {
	size := l.size
	p, _ := l.First()
	for size > 0 {
		if l.equal(p.data(), e) <= 1 {
			break
		}
		p = p.Pred()
	}
	return p
}

// 唯一化
func (l *list[T]) Deduplicate() int {
	if l.size < 2 { //重复元素的个数
		return 0
	}
	oldSize := l.size
	root, _ := l.First()           //拿到第一个元素
	r := 1                         //从的一个元素开始
	for root.Succ() != l.trailer { //尾节点之前
		q := l.findNode(root.data(), r, root) //范围内查找指定元素
		if q != nil {
			l.Remove(q)
		} else {
			r++
		}
		root = root.Succ()
	}
	return oldSize - l.size
}

func (l *list[T]) Uniquify() int {
	if l.size < 2 {
		return 0
	}
	oldSize := l.size
	root, _ := l.First()
	q := root.Succ()
	for l.trailer != root.Succ() {
		if l.equal(root.data(), q.data()) > 0 {
			root = q
		} else {
			l.Remove(q)
		}
	}
	return oldSize - l.size
}

func (l *list[T]) Traverse(f func(e T)) {
	//TODO implement me
	panic("implement me")
}

func NewList[T any]() List[T] {
	header := NewNode(*new(T)) //创建个默认值，不影响啥
	trailer := NewNode(*new(T))
	header.InsertAsSucc(trailer)
	trailer.InsertAsPred(header)
	return &list[T]{
		size:    0,
		header:  header,
		trailer: trailer,
	}
}

func (l *list[T]) Clear() {
	l.size = 0
	l.header.InsertAsSucc(l.trailer)
	l.trailer.InsertAsPred(l.header)
}

func (l *list[T]) First() (ListNode[T], error) {
	if l.size == 0 {
		return l.header, errors.New("当前列表为空")
	}
	return l.header.Succ(), nil
}

func (l *list[T]) Get(index int) (ListNode[T], error) {
	if index > l.size || index < 0 {
		return l.header, errors.New("数组越界")
	}
	root := l.header
	for i := 0; i < index; i++ {
		root = root.Pred()
	}
	return root, nil
}

func (l *list[T]) Last() (ListNode[T], error) {
	if l.size == 0 {
		return l.header, errors.New("当前列表为空")
	}
	return l.trailer.Pred(), nil
}

func (l *list[T]) InsertAsFirst(val T) {
	node := NewNode[T](val) //新节点
	old := l.header.Succ()  //旧的后驱
	l.header.InsertAsSucc(node)
	node.InsertAsSucc(old)
	old.InsertAsPred(node)
}

func (l *list[T]) InsertAsLast(val T) {
	node := NewNode[T](val) //新节点
	old := l.trailer.Pred() //旧的前驱
	old.InsertAsSucc(node)
	node.InsertAsPred(old)
	l.trailer.InsertAsPred(node)
}

func (l *list[T]) InsertBefore(p ListNode[T], e T) {
	node := NewNode[T](e) //新节点
	p.InsertAsSucc(node)
	node.InsertAsPred(p)
}

func (l *list[T]) InsertAfter(p ListNode[T], e T) {
	node := NewNode[T](e) //新节点
	p.InsertAsPred(node)
	node.InsertAsSucc(p)
}

func (l *list[T]) Remove(p ListNode[T]) T {
	val := p.data()
	pred := p.Pred()
	succ := p.Succ()
	pred.InsertAsSucc(succ)
	succ.InsertAsPred(succ)
	return val
}

// 查找
func (l *list[T]) findNode(data T, r int, root ListNode[T]) ListNode[T] {
	for r > 0 {
		pred := root.Pred()
		if l.equal(pred.data(), data) == 0 {
			//结束，返回
			return pred
		}
		r--
	}
	return nil
}
