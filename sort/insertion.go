package sort

import (
	"sort"
)

//插入排序，兼容标准库排序接口
func Insertion(data sort.Interface) {
	for i:=1; i < data.Len(); i++ {
		for x:=i-1; x >= 0; x-- {
			if data.Less(x, x+1) {
				break;
			} else {
				data.Swap(x, x+1)
			}
		}
	}
}
