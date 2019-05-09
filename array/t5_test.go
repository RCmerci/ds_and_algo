package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t5(arr []int, key int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if arr[mid] == key {
			return mid
		}
		if arr[mid] < arr[l] { // 3,4,5,[1],2
			if key <= arr[r-1] && key > arr[mid] {
				l = mid + 1
			} else {
				r = mid
			}
		} else { // 3,4,[5],1,2
			if key >= arr[l] && key < arr[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

func Test_t5(t *testing.T) {
	var cases = []struct {
		test   []int
		key    int
		expect int
	}{
		{[]int{3, 4, 5, 1, 2}, 1, 3},
		{[]int{3, 4, 5, 1, 2}, 6, -1},
		{[]int{4, 5, 6, 1, 3}, 2, -1},
		{[]int{3, 4, 5, 1, 2}, 3, 0},
		{[]int{5, 6, 7, 8, 9, 10, 1, 2, 3}, 3, 8},
		{[]int{5, 6, 7, 8, 9, 10, 1, 2, 3}, 30, -1},
		{[]int{}, 30, -1},
	}
	for i := range cases {
		assert.Equal(t, cases[i].expect, t5(cases[i].test, cases[i].key))
	}

}
