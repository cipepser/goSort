package MySortCahngeN 

import (
	"testing"
	"math/rand"
	"time"
	"sort"
)

func BenchmarkMySortOffset(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		rand.Perm(len(a))
	}
}


func BenchmarkMySortRandomizedDataGoPackageTimer(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		b.StartTimer()
		sort.Ints(tmp)
		b.StopTimer()
	}
}

func BenchmarkMySortRandomizedData(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		
		MySort(tmp)
	}
}

func BenchmarkMySortRandomizedDataTimer(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		b.StartTimer()
		MySort(tmp)
		b.StopTimer()
	}
}
