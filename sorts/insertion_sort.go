package sorts

func InsertionSort(ary []int) []int {
	for i := 1; i < len(ary); i++ {

		for j := i - 1; j >= 0; j-- {
			if ary[j+1] > ary[j] {
				break
			}

			ary[j], ary[j+1] = ary[j+1], ary[j]
		}
	}
	return ary
}
