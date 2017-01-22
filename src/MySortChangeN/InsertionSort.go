package MySortCahngeN

func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i - j - 1] > a[i - j] {
				a[i - j - 1], a[i - j] = a[i - j], a[i - j - 1]
			} else {
				break
			}
		}
	}
}
