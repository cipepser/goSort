package MySortCahngeN

import (
	"testing"
	"math/rand"
	"time"
)

// 乱数生成のoffset
func BenchmarkBaseRandom(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		rand.Perm(len(a))
	}
}
