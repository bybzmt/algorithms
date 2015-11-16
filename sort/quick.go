package sort

import (
	"sort"
)

//快速排序，兼容标准库排序接口
func Quick(data sort.Interface) {
	quick_do(data, 0, data.Len()-1)
}

//快速排序， start和end是指需要操作的下标区间
func quick_do(data sort.Interface, start, end int) {
	if end > start {
		point := quick_partition(data, start, end)
		quick_do(data, start, point-1)
		quick_do(data, point+1, end)
	}
}

//快速排序分区
func quick_partition(data sort.Interface, start, end int) int {
	point := start

	for i:=start; i<end; i++ {
		if data.Less(i,end) {
			if i != point {
				data.Swap(i, point)
			}
			point++
		}
	}

	data.Swap(point, end)
	return point
}
