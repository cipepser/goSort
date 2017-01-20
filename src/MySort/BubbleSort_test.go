package MySort 

import (
	"testing"
	"reflect"
	"sort"
	// "fmt"
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
	for i:= 0; i <b.N; i++ {
		tmp := make([]int, len(a))
		copy(tmp, a)
		// b.StopTimer()
		// // 初期化処理
		// b.StartTimer()

		BubbleSort(a)
	}
}
