package sort

import (
	"testing"
	"sort"
)

func TestQuick(t *testing.T) {
	d := init_testdata(20)
	Quick(d)
	if !sort.IsSorted(d) {
		t.Error("Quick Sort Error:", d)
		return
	}
}
