package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t6(arr []int, x int) bool {
	if len(arr) < 2 {
		return false
	}
	if len(arr) == 2 {
		return arr[0]+arr[1] == x
	}
	// find pivot
	l, r := 0, len(arr)
	pivot := -1
	for l < r {
		mid := (l + r) / 2
		if arr[mid] > arr[mid+1] {
			pivot = mid
			break
		}
		if arr[mid] < arr[l] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	min, max := 0, len(arr)-1
	if pivot != -1 {
		min, max = pivot+1, pivot
	}

	for min != max {
		if arr[min]+arr[max] == x {
			return true
		}
		if arr[min]+arr[max] < x {
			min = (min + 1) % len(arr)
		} else {
			max = (len(arr) + max - 1) % len(arr)
		}
	}
	return false

}

func Test_t6(t *testing.T) {
	var cases = []struct {
		test   []int
		x      int
		expect bool
	}{
		{[]int{11, 15, 6, 8, 9, 10}, 16, true},
		{[]int{11, 15, 26, 38, 9, 10}, 35, true},
		{[]int{11, 15, 26, 38, 9, 10}, 45, false},
		{[]int{0, 1}, 1, true},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t6(cases[i].test, cases[i].x))
	}
}
