// Package bubble-sort provides ...
package sorts

func BubbleSort(ary []int) []int {
	for n := len(ary); n >= 0; n-- {
		for i := 1; i < n; i++ {
			if ary[i-1] > ary[i] {
				ary[i-1], ary[i] = ary[i], ary[i-1]
			}
		}
	}

	return ary
}
