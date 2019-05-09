package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t9(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	l, r := 0, len(arr)
	mid := -1
	for l < r {
		mid = (l + r) / 2
		if mid+1 < len(arr) && arr[mid] > arr[mid+1] {
			return mid + 1
		}
		if arr[mid] > arr[l] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return (mid + 1) % len(arr)
}

func Test_t9(t *testing.T) {
	var cases = []struct {
		test   []int
		expect int
	}{
		{[]int{15, 18, 2, 3, 6, 12}, 2},
		{[]int{7, 9, 11, 12, 5}, 4},
		{[]int{7, 9, 11, 12, 15}, 0},
		{[]int{1, 2}, 0},
		{[]int{2, 1}, 1},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t9(cases[i].test))
	}

}
