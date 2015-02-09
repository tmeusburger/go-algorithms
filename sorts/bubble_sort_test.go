// Package sorts provides ...
package sorts

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testIntegerSortFunction(t, BubbleSort)
}
