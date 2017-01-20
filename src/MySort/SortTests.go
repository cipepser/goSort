package MySort

import (
	// "math/rand"
	// "time"
	// "fmt"
)

// テストパターン
var SortTests = [][]int {
	{2, 1, 3, 5, 4},
	{1, 2, 3, 4, 5},
	{1, 3, 4, 5, 2, 1}, 
	{1, 0, 2},
	{1, 1, 1, 1, 1, 1, 1, 1},
}

// ベンチマーク
// var a = []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var a = []int {0, 3, 1, 2, 3, 4, 5, 1, 9, 0}
// var a []int
// 
// func init()  {
// 	rand.Seed(time.Now().UnixNano())
// 	a = rand.Perm(10)	
// 	fmt.Println(a)
// }
