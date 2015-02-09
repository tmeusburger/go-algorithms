package sorts

func MergeSort(ary []int) []int {
	length := len(ary)

	if length == 1 {
		return ary
	}

	mid := length / 2
	left := MergeSort(ary[:mid])
	right := MergeSort(ary[mid:])

	sorted := merge(left, right)
	return sorted
}

func merge(l, r []int) []int {
	result := []int{}

	var i, j int = 0, 0
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			result = append(result, l[i])
			i++
		} else {
			result = append(result, r[j])
			j++
		}
	}

	for ; i < len(l); i++ {
		result = append(result, l[i])
	}

	for ; j < len(r); j++ {
		result = append(result, r[j])
	}

	return result
}
