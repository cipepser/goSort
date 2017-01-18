package MySort 

import (
	"testing"
	"reflect"
	"sort"
)

func TestInsertionSort(t *testing.T) {
	for _, actual := range SortTests {
		expected := make([]int, len(actual))
		copy(expected, actual)
		
		sort.Ints(expected)
		InsertionSort(actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		InsertionSort(a)
	}
}
