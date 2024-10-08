package list

import (
	"github.com/Diana-Fox/go_data_structure/data_structure/iterator"
)

// List 定义一个list,其实完全不用返回error，而是应该报错
type List[T any] interface {
	Size() int
	Get(index int) (T, error)
	Set(index int, v T) error
	Insert(index int, v T) error
	Append(v T)
	Clear()
	Delete(i int) (T, error)
	String() string
	Iterator() iterator.Iterator[T]
}
