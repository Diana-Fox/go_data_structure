package application_info

import (
	"fmt"
	"github.com/Diana-Fox/go_data_structure/data_structure/stack"
	"testing"
)

// 用别名
type Integer int

// 使用栈模拟递归，low版本
func TestSimulateRecursion(t *testing.T) {
	arrayStack := stack.NewArrayStack[Integer]()
	arrayStack.Push(5)
	var count Integer = 0
	for !arrayStack.IsEmpty() {
		pop, err := arrayStack.Pop()
		if err != nil {
			t.Error(err)
		}
		count += pop
		if pop > 0 {
			arrayStack.Push(pop - 1)
		}
	}
	fmt.Println(count)
}

// 非波拉契数
func TestFoc(t *testing.T) {
	//recursion := FABRecursion(6)
	recursion, err := FABStack(6)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(recursion)
}

// 递归版的非波拉契数
func FABRecursion(num int) int {
	if num <= 2 {
		return 1
	}
	return FABRecursion(num-1) + FABRecursion(num-2)
}

// 栈实现非波拉契数
func FABStack(num int) (int, error) {
	stack := stack.NewArrayStack[Integer]()
	stack.Push(Integer(num))
	var count Integer = 0
	for !stack.IsEmpty() {
		pop, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			return int(count), err
		}
		if pop > 2 {
			stack.Push(pop - 2)
			stack.Push(pop - 1)
		} else {
			count += 1
		}
	}
	return int(count), nil
}
