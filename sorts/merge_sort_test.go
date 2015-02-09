package sorts

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	testIntegerSortFunction(t, MergeSort)
}
