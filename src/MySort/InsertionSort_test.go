package MySort 

import (
	"testing"
	"reflect"
	"sort"
	"math/rand"
	"time"
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
		tmp := make([]int, len(a))
		copy(tmp, a)
		
		InsertionSort(tmp)
	}
}

func BenchmarkInsertionSorted(b *testing.B) {
	tmp := make([]int, len(SortedData))

	for i:= 0; i <b.N; i++ {
		copy(tmp, SortedData)

		InsertionSort(tmp)
	}
}

func BenchmarkInsertionSortRandomizedData(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		
		InsertionSort(tmp)
	}
}
