package sort

import (
	"testing"
	"sort"
)

func TestInsertion(t *testing.T) {
	d := init_testdata(20)
	Insertion(d)
	if !sort.IsSorted(d) {
		t.Error("Insertion Sort Error:", d)
		return
	}

}
