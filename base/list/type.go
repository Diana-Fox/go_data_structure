package list

type List[T any] interface {
	Size() int                          //列表大小
	First() (ListNode[T], error)        //头节点
	Get(index int) (ListNode[T], error) //获取下表为index的元素
	Last() (ListNode[T], error)         //尾节点
	Clear()                             //清空
	InsertAsFirst(val T)                //插入为头节点
	InsertAsLast(val T)                 //插入为尾节点
	InsertBefore(p ListNode[T], e T)    //插入指定节点后
	InsertAfter(p ListNode[T], e T)     //插入指定节点前
	Remove(p ListNode[T]) T             //移除节点
	Disordered()                        //判断所有节点是否已经按非降序排列
	Sort()                              //排序，使节点按非降序排列
	Find(e T) ListNode[T]               //查找目标元素
	Search(e T) ListNode[T]             //搜索，返回不大于e且index最大的节点
	Deduplicate() int                   //删除重复节点
	Uniquify() int                      //唯一化
	Traverse(func(e T))                 //遍历
}
