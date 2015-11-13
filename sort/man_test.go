package sort

import (
	"math/rand"
	"testing"
)

type testdata []int

func (d testdata) Len() int {
	return len(d)
}

func (d testdata) Less(i, j int) bool {
	return d[i] < d[j]
}

func (d testdata) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func init_testdata(n int) testdata {
	d := make(testdata, n)

	for i:=0; i < n; i++ {
		d[i] = rand.Intn(n*2)
	}

	return d
}


func TestMain(m *testing.M) {
	m.Run()
}
