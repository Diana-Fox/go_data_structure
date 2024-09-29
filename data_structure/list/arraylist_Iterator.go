package list

// ArrayList的迭代器
type ArrayListIterator[T any] struct {
	list         *ArrayList[T]
	currentIndex int
}

func (a *ArrayListIterator[T]) HasNext() bool {
	if a.currentIndex >= a.list.Size() {
		return false
	}
	return true
}

func (a *ArrayListIterator[T]) Next() (T, error) {
	value, err := a.list.Get(a.currentIndex)
	if err != nil {
		return value, err
	}
	a.currentIndex++
	return value, nil
}

func (a *ArrayListIterator[T]) Remove() error {
	_, err := a.list.Delete(a.currentIndex)
	return err
}

func (a *ArrayListIterator[T]) GetIndex() int {
	return a.currentIndex
}
