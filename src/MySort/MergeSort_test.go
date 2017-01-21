package MySort 

import (
	"testing"
	"reflect"
	"sort"
)

func TestMergeSort(t *testing.T) {
	for _, actual := range SortTests {
		expected := make([]int, len(actual))
		copy(expected, actual)
		
		sort.Ints(expected)
		MergeSort(actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		tmp := make([]int, len(a))
		copy(tmp, a)

		MergeSort(tmp)
	}
}

func BenchmarkMergeSorted(b *testing.B) {
	tmp := make([]int, len(SortedData))

	for i:= 0; i <b.N; i++ {
		copy(tmp, SortedData)

		MergeSort(tmp)
	}
}

func BenchmarkMergeSortRandomizedData(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		
		MergeSort(tmp)
	}
}
