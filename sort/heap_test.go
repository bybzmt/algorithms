package sort

import (
	"testing"
	"sort"
)

func TestHeapSort(t *testing.T) {
	d := init_testdata(21)
	Heap(d)
	if !sort.IsSorted(d) {
		t.Error("Heap Sort Error:", d)
		return
	}

}
