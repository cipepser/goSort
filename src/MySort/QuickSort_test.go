package MySort 

import (
	"testing"
	"reflect"
	"sort"
	"math/rand"
	"time"
)

func TestQuickSort(t *testing.T) {
	for _, actual := range SortTests {
		expected := make([]int, len(actual))
		copy(expected, actual)
		
		sort.Ints(expected)
		QuickSort(actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
		}
	}
}

func BenchmarkMQuickSort(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		tmp := make([]int, len(a))
		copy(tmp, a)

		QuickSort(tmp)
	}
}

func BenchmarkQuickSorted(b *testing.B) {
	tmp := make([]int, len(SortedData))

	for i:= 0; i <b.N; i++ {
		copy(tmp, SortedData)

		QuickSort(tmp)
	}
}

func BenchmarkQuickSortRandomizedData(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		
		QuickSort(tmp)
	}
}
