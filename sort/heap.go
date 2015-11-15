package sort

import (
	"sort"
)

//堆排序，兼容标准库排序接口
func Heap(data sort.Interface) {
	h := new(heap).init(data)

	for i:=data.Len()-1; i>0; i-- {
		data.Swap(0, i)
		h.size--
		h.heapify(0)
	}
}

type heap struct {
	size int
	data sort.Interface
}

func heap_parent(i int) int {
	return (i-1)/2
}

func heap_left(i int) int {
	return (i+1)*2-1
}

func heap_right(i int) int {
	return (i+1)*2
}

func (h *heap) init(data sort.Interface) *heap {
	h.data = data
	h.size = data.Len()

	for i:=data.Len()/2; i>=0; i-- {
		h.heapify(i)
	}

	return h
}

func (h *heap) heapify(i int) {
	l := heap_left(i)
	r := heap_right(i)
	large := i

	if l < h.size && h.data.Less(i, l) {
		large = l
	}

	if r < h.size && h.data.Less(large, r) {
		large = r
	}

	if large != i {
		h.data.Swap(i, large)
		h.heapify(large)
	}
}
