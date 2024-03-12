package list

import (
	"fmt"
	"testing"
)

func TestArrayListIterator(t *testing.T) {
	list := NewArrayList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}
	for it := list.Iterator(); it.HashNext(); {
		next, _ := it.Next()
		fmt.Println(next)
	}
}
