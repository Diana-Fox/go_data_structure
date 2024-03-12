package list

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}
	list.Insert(5, 50)
	fmt.Print(list)
}
