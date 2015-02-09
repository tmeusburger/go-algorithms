package sorts

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type intSortingFunction func([]int) []int

func testIntSort(t *testing.T, sf intSortingFunction) {
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8}

	Convey("When provided an unorder slice", t, func() {
		Convey("Of even length", func() {
			sorted := []int{1, 2, 3, 4, 5, 6, 7, 8}
			unsorted := []int{1, 2, 5, 3, 7, 4, 6, 8}

			Convey("Returns a sorted slice", func() {
				So(sf(unsorted), ShouldResemble, sorted)
			})
		})

		Convey("Of odd length", func() {
			sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			unsorted := []int{9, 1, 2, 5, 3, 7, 4, 6, 8}

			Convey("Returns a sorted slice", func() {
				So(sf(unsorted), ShouldResemble, sorted)
			})
		})
	})

	Convey("When provided an ordered slice", t, func() {
		sorted = []int{1, 2, 3, 4, 5, 6, 7, 8}

		Convey("Returns a sorted slice", func() {
			So(sf(sorted), ShouldResemble, sorted)
		})
	})

}
