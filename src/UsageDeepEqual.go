package main

import (
	"fmt"
	"reflect"
)

func main()  {
	a := []int{1, 2, 3}
	b := []int{2, 1, 3} // a != b
	
	answer := reflect.DeepEqual(a, b) // 要素がひとつ異なる
	fmt.Println(answer) // false
	answer = reflect.DeepEqual(a, a) // 自分自身なので等しい
	fmt.Println(answer) // true
}