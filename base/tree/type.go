package tree

// 长子兄弟表示法
type TheTree[T any] interface {
	Root()                //根节点
	Parent()              //父节点
	FirstChild()          //长子
	NextSibling()         //兄弟
	Insert(i int, val T)  //将val作为第i个孩子插入
	Remove(i int)         //删除第i个节点
	Tracerse(func(val T)) //遍历节点
}

// 二叉树
type BinTree[T any] interface {
	InsertAsRC(node BinNode[T], val T) *BinNode[T]
	TravLevel(node Node[T]) //层次遍历
	TravPre(node Node[T])   //先序遍历
	TravIn(node Node[T])    //中序遍历
	TravPost(node Node[T])  //后序遍历
}
