package sorts

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	testIntegerSortFunction(t, InsertionSort)
}
