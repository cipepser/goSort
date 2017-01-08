package main

import (
	"fmt"
)

// 中間値を返す
func Med3(x, y, z int) int {
	if x < y {
		if y < z {
			return y
		} else if x < z {
			return z
		} else {
			return x
		}
	} else {
		if x < z {
			return x
		} else if y < z {
			return z
		} else {
			return y
		}
	}
}

func QuickSort(a []int) []int {
	pivot := Med3(a[0], a[len(a) / 2], a[len(a) - 1])
	left := 0
	right := len(a) - 1
// tmp := 0
	for left < right {
		for a[left] < pivot {
			left++
		}
	
		for a[right] > pivot {
			// if left == right {
			// 	break
			// }
			right--
		}
		
		// if left >= right {
		// 	break
		// }

		fmt.Println("-------------------")
		fmt.Printf("a[left]: %d\n", a[left])
		fmt.Printf("left: %d\n", left)
		fmt.Printf("a[right]: %d\n", a[right])
		fmt.Printf("right: %d\n", right)
		fmt.Printf("pivot: %d\n", pivot)
		fmt.Printf("a: %v\n", a)

		a[left], a[right] = a[right], a[left]
		// left++
		// right--

		fmt.Printf("a: %v\n", a)
		
		// tmp++
		// if tmp > 3 {
		// 	break
		// }

		// left++
		// right--
	}
	fmt.Println("*******************")
	fmt.Println(a)
	a1 := a[:left]
	fmt.Printf("a1: %v\n", a1)
	if len(a1) > 1 {
		a1 = QuickSort(a1)
	}

	a2 := a[right+1:]
	fmt.Printf("a2: %v\n", a2)
	if len(a2) > 1 {
		a2 = QuickSort(a2)
	}
	
	cnt := 0
	for _, v := range a1 {
		a[cnt] = v
		cnt++
	}
	a[cnt] = pivot
	cnt++
	for _, v := range a2 {
		a[cnt] = v
		cnt++
	}
	
	return a
}

func main()  {
	a := []int{2, 4, 5, 1, 3}
	// a := []int{2, 4, 1, 3, 1}
	fmt.Println(QuickSort(a))
}