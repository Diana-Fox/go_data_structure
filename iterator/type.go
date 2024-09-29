package iterator

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
