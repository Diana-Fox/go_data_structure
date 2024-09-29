package list

type List[T any] interface {
	Size() int
	Get(index int) (T, error)
	Set(index int, v T) error
	Insert(index int, v T) error
	Append(v T)
	Clear()
	Delete(i int) error
	String() string
	Iterator() Iterator[T]
}

// Iterator 实际的迭代器
type Iterator[T any] interface {
	HasNext() bool
	Next() (T, error)
	Remove() error //删除
	GetIndex() int
}

// Iterable 迭代器功能接口
type Iterable[T any] interface {
	Iterator() Iterator[T] //迭代器接口
}
