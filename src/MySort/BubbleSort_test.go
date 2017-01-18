package MySort 

import (
	"testing"
	"reflect"
	"sort"
)

func TestBubbleSort(t *testing.T) {
	for _, actual := range SortTests {
		expected := make([]int, len(actual))
		copy(expected, actual)
		
		sort.Ints(expected)
		BubbleSort(actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	a := []int{0, 3, 1, 2, 3, 4, 5, 1, 9, 0}
	// b.ResetTimer()
	for i:= 0; i <b.N; i++ {
		BubbleSort(a)
	}
}
