package MySort 

import (
	"testing"
	"reflect"
	"sort"
)

func TestHeapSort(t *testing.T) {
	for _, actual := range SortTests {
		expected := make([]int, len(actual))
		copy(expected, actual)
		
		sort.Ints(expected)
		HeapSort(actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
		}
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		tmp := make([]int, len(a))
		copy(tmp, a)

		HeapSort(tmp)
	}
}

func BenchmarkHeapSorted(b *testing.B) {
	tmp := make([]int, len(SortedData))

	for i:= 0; i <b.N; i++ {
		copy(tmp, SortedData)

		HeapSort(tmp)
	}
}

func BenchmarkHeapSortRandomizedData(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		
		HeapSort(tmp)
	}
}
