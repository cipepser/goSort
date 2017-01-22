package MySortCahngeN 

import (
	"testing"
	"math/rand"
	"time"
	"sort"
)

func BenchmarkMySortRandomizedDataGoPackage(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		
		sort.Ints(tmp)
	}
}


func BenchmarkMySortRandomizedData(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(a))
		
		MySort(tmp)
	}
}
