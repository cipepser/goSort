package main

import (
	"fmt"
	"sort"
)

func main()  {
	a := []int{2, 3, 1, 2, 5}
	
	sort.Ints(a)
	fmt.Println(a)
}