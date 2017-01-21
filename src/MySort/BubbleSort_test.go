package MySort 

import (
	"testing"
	"reflect"
	"sort"
	"math/rand"
	"time"
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
	tmp := make([]int, len(a))

	for i:= 0; i <b.N; i++ {
		copy(tmp, a)

		BubbleSort(tmp)
	}
}

func BenchmarkBubbleSorted(b *testing.B) {
	tmp := make([]int, len(SortedData))

	for i:= 0; i <b.N; i++ {
		copy(tmp, SortedData)

		BubbleSort(tmp)
	}
}

func BenchmarkBubbleSortRandomizedData(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		
		BubbleSort(tmp)
	}
}
