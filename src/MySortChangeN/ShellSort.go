package MySortCahngeN

func CalcInterval(n int) int {
	h := 1
	
	for h < n {
		h = 3 * h + 1
	}
	
	h = (h - 1) / 3
	
	return h
}

func ShellSort(a []int) {
	h := CalcInterval(len(a))
	
	for h > 1 {
		for i := 0; i < h; i++ {
			// hずつ飛ばしたグループを作る
			b := make([]int, len(a) / h + 1)
			cnt := 0
			for j := 0; j < len(a) / h + 1; j++ {
				if i + h * j < len(a){
					b[j] = a[i + h * j]
					cnt++
				}
			}
			
			// 抜き出したグループに対して挿入ソート
			c := b[:cnt]
			InsertionSort(c)
			
			// ソート済みのものを代入
			for j := 0; j < len(c); j++ {
				if i + h * j < len(a){
					a[i + h * j] = c[j]
				}
			}
			
		}
		h = (h - 1) / 3
	}
	
	InsertionSort(a)
}
