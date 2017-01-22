package MySortCahngeN
 
func Min(a []int) (idx, n int) {	
	n = a[0]
	idx = 0
	for i, v := range a {
		if n > v {
			n = v
			idx = i
		}
	}
	
	return
}

func SelectionSort(a []int) {
	for i, _ := range a {
		idx, _ := Min(a[i:len(a)])
		a[i], a[i + idx] = a[i + idx], a[i]
	}
}
