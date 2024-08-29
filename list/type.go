package list

type List[T any] interface {
	Size() int
	Get(i int) (T, error)
	Set(i int, v T) error
	Insert(i int, v T) error
	Append(v T)
	Clear()
	Delete(i int) error
	String() string
}
