package MySortCahngeN 

import (
	"testing"
	"math/rand"
	"time"
)

func BenchmarkSelectionSorted(b *testing.B) {
	b.StopTimer()
	for i, _ := range a {
		a[i] = i
	}
	tmp := a
	b.StartTimer()

	for i:= 0; i <b.N; i++ {
		SelectionSort(tmp)
	}
}

func BenchmarkSelectionSortRandomizedData(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		
		SelectionSort(tmp)
	}
}
