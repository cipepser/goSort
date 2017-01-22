package MySort

import (
	"testing"
	"math/rand"
	"time"
)

// コピーのoffset
func BenchmarkBaseCopy(b *testing.B) {
	tmp := make([]int, len(a))

	for i:= 0; i <b.N; i++ {
		copy(tmp, a)
	}
}

// 乱数生成のoffset
func BenchmarkBaseRandom(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		rand.Perm(len(a))
	}
}
