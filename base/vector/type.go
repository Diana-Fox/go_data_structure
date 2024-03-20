package vector

type Vector[T any] interface {
	CopyFrom(vector []T, lo, hi int)             //通过数组复制
	CopyFromVector(vector Vector[T], lo, hi int) //通过数组复制
	Get(index int) (T, error)                    //获取下表为index的元素
	Put(val T)                                   //追加元素
	Insert(index int, val T)                     //将val插入到下标为index的位置
	Remove(lo, hi int) int                       //移除元素
	Find(val T, lo, hi int) int                  //区间内元素
	Eques(val1 T, val2 T) bool                   //元素大小比较，理论上泛型不支持比较，但是具体实现类可以去实现比较
	Traverse(func(val T))                        //遍历操作
	Disordered() int                             //逆序对的个数
	Uniquify() int                               //有序向量的唯一化,返回删除元素的个数
	Seaech(val T, lo, hi int) (int, error)       //有序向量的查找
	Sort(lo, hi int)                             //排序算法
	Size() int                                   //大小
}
