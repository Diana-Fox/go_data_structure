package vector

import (
	"errors"
	"math/rand/v2"
)

type vector[T any] struct {
	data     []T //go没有传统意义上的数组，只有这个切片，切片其实跟vector机制相似
	size     int
	capacity int
	eques    func(v1 T, v2 T) int //比较的方法具体实现
}

func (v *vector[T]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (v *vector[T]) Sort(lo, hi int) {
	switch rand.Int() % 5 {
	case 1:
		v.bubbleSort(lo, hi)
		break
	case 2:
		v.selectionSort(lo, hi)
		break
	case 3:
		v.mergeSort(lo, hi)
		break
	case 4:
		v.heapSort(lo, hi)
		break
	default:
		v.quickSort(lo, hi)
	}
}

// 元素交换
func (v *vector[T]) swap(lo, hi int) {
	temp := v.data[lo]
	v.data[lo] = v.data[hi]
	v.data[hi] = temp
}
func (v *vector[T]) bubbleSort(lo, hi int) {
	for lo < (hi) {
		//将最后一次发生交换的位置设置为hi
		hi = v.bubble(lo, hi)
	}
}

// 智能的发现最后一次交换的位置，后面都为有序，直接不进行迭代
func (v *vector[T]) bubble(lo, hi int) int {
	last := lo
	lo++
	for lo < hi {
		if v.eques(v.data[lo-1], v.data[lo]) > 0 { //找到了逆序对
			last = lo
			v.swap(lo-1, lo) //去交换
		}
		lo++
	}
	return last
}

// 稍微没那么智能的版本
func (v *vector[T]) bubbleSortA(lo, hi int) {
	for !v.bubbleA(lo, hi) {
		hi--
	}
}

// 只能发现前段排序完成，提前结束，不能发现后段
func (v *vector[T]) bubbleA(lo, hi int) bool {
	sorted := true
	lo++
	for lo < hi {
		if v.eques(v.data[lo-1], v.data[lo]) > 0 { //找到了逆序对
			sorted = false
			v.swap(lo-1, lo) //去交换
		}
		lo++
	}
	return sorted
}
func (v *vector[T]) selectionSort(lo, hi int) {

}
func (v *vector[T]) mergeSort(lo, hi int) {
	if hi-lo < 2 {
		return
	}
	mi := (hi + lo) >> 1 //找到中点
	v.mergeSort(lo, mi)  //左半段排序
	v.mergeSort(mi, hi)  //右半段排序
	v.merge(lo, mi, hi)  //合并左右
}

// 排序
func (v *vector[T]) merge(lo, mi, hi int) {
	lb := mi - lo         //左边的长度
	b := make([]T, 0, lb) //左边向量
	for i := 0; i < mi-lo; i++ {
		b[i] = v.data[i]
	}
	lc := hi - mi //右边的长度
	i := 0
	j := 0
	k := 0
	for (j < lb) || (k < lc) { //说明两边向量还有元素
		//找出更小的那个
		if (j < lb) && (lc <= k || (v.eques(b[j], v.data[k]) < 1)) {
			v.data[i] = b[j]
			i++
			j++
		}
		if (k < lc) && (lb <= j || (v.eques(v.data[j], v.data[k]) < 1)) {
			v.data[i] = b[k]
			i++
			k++
		}
	}

}
func (v *vector[T]) heapSort(lo, hi int) {

}
func (v *vector[T]) quickSort(lo, hi int) {

}

// 有序向量的查找,返回不大于val的最后一个元素
func (v *vector[T]) Seaech(val T, lo, hi int) (int, error) {
	if lo > hi {
		return 0, errors.New("查找范围异常")
	}
	if rand.Int()%2 == 0 {
		return v.binSearch(val, lo, hi), nil
	}
	return v.FibonacciSearch(val, lo, hi), nil
}

// 二分查找
func (v *vector[T]) binSearch(val T, lo, hi int) int {
	for lo < hi {
		mi := (lo + hi) >> 1
		if v.eques(val, v.data[mi]) < 0 {
			hi = mi
		} else {
			lo = mi
		}
	}
	lo--
	return lo
}

// 二分查找，但是没有严格返回不大于val的值
func (v *vector[T]) binSearchV1(val T, lo, hi int) int {
	for 1 < hi-lo {
		mi := (lo + hi) >> 1
		if v.eques(val, v.data[mi]) > 0 {
			lo = mi
		} else {
			hi = mi
		}
	}
	if v.eques(val, v.data[lo]) == 0 {
		return lo
	}
	return -1
}

// 斐波拉契查找
func (v *vector[T]) FibonacciSearch(val T, lo, hi int) int {
	//先通过hi-lo获得一个fib数
	//根据fib迭代找到黄金分割点
	panic("暂时不实现了")
}
func NewVector[T any](eques func(v1 T, v2 T) int) Vector[T] {
	return &vector[T]{
		data:     make([]T, 0, 10),
		capacity: 10,
		size:     0,
		eques:    eques,
	}
}

func (v *vector[T]) Uniquify() int {
	i := 0
	j := 1
	for j < v.size {
		if v.eques(v.data[i], v.data[j]) != 0 {
			i++
			//互异
			v.data[i] = v.data[j]
		}
		j++
	}
	v.size = i + 1
	v.shrink()
	return j - i
}

func (v *vector[T]) Disordered() int {
	n := 0
	for i := 1; i < v.size; i++ {
		if v.eques(v.data[i-1], v.data[i]) > 0 {
			n++
		}
	}
	return n
}

func (v *vector[T]) Traverse(f func(val T)) {
	for i := 0; i < v.size; i++ {
		f(v.data[i])
	}
}

func (v *vector[T]) Find(val T, lo, hi int) int {
	for lo < hi && v.eques(v.data[hi], val) != 0 {
		hi--
	}
	return hi
}

// 比较的方法
func (v *vector[T]) Eques(val1 T, val2 T) bool {
	return v.eques(val1, val2) == 0
}

func (v *vector[T]) Remove(lo, hi int) int {
	if lo == hi {
		return 0
	}
	//
	for hi < v.size {
		v.data[lo] = v.data[hi]
		lo++
		hi++
	}
	v.size = lo
	v.shrink()
	return hi - lo
}

func (v *vector[T]) Put(val T) {
	v.expand()
	v.data[v.size] = val
	v.size++
}

func (v *vector[T]) Insert(index int, val T) {
	v.expand()
	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = val
}

func (v *vector[T]) Get(index int) (T, error) {
	if index > v.size {
		return *new(T), errors.New("数组越界")
	}
	return v.data[index], nil
}
func (v *vector[T]) CopyFrom(vector []T, lo, hi int) {
	v.data = make([]T, 0, hi-lo*2) //
	v.size = 0
	for lo < hi {
		v.data[v.size] = vector[lo]
		v.size++ //不支持在[]直接++，拆出来
		lo++
	}
}

func (v *vector[T]) CopyFromVector(vector Vector[T], lo, hi int) {

}

// 扩容策略
func (v *vector[T]) expand() {
	if v.size < v.capacity {
		return
	}
	oldData := v.data
	v.data = make([]T, 0, v.capacity<<2)
	for i := 0; i < v.size; i++ {
		v.data[i] = oldData[i]
	}
}

// 缩容策略
func (v *vector[T]) shrink() {
	if v.size < v.capacity/2 {

	}
}
