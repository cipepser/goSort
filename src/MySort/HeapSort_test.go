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
	a := []int{0, 3, 1, 2, 3, 4, 5, 1, 9, 0}
	for i:= 0; i <b.N; i++ {
		HeapSort(a)
	}
}
