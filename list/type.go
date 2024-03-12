package list

type List[T any] interface {
	Size() int                          //总个数
	Get(index int) (T, error)           //获取第几个元素
	Set(index int, value T) error       //修改元素
	Insert(index int, newValue T) error //插入元素
	Append(newValue T)                  //末尾追加元素
	Clear()                             //清空
	Delete(index int) error             //删除第几个元素
	String() string                     //返回字符串
}

// -------------------------------
// 迭代器
type Iterator[T any] interface {
	HashNext() bool   //是否有下一个
	Next() (T, error) //下一个
	Remove()          //删除
}

// 迭代
type Iterable[T any] interface {
	Iterator() Iterator[T] //构造必须初始化这个接口
}
