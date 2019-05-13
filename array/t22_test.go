package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t22(arr []int) []int {
	l, r := 0, 0
	for l < len(arr) {
		if l%2 == 0 {
			if arr[l] >= 0 {
				for r < len(arr) && arr[r] >= 0 {
					r++
				}
				if r == len(arr) {
					return arr
				}
				tmp := arr[r]
				for i := r; i > l; i-- {
					arr[i] = arr[i-1]
				}
				arr[l] = tmp
			}
		} else {
			if arr[l] < 0 {
				for r < len(arr) && arr[r] < 0 {
					r++
				}
				if r == len(arr) {
					return arr
				}
				tmp := arr[r]
				for i := r; i > l; i-- {
					arr[i] = arr[i-1]
				}
				arr[l] = tmp
			}
		}
		l++
		r = l
	}
	return arr
}

func Test_t22(t *testing.T) {
	var cases = []struct {
		test   []int
		expect []int
	}{
		{[]int{1, 2, 3, -4, -1, 4}, []int{-4, 1, -1, 2, 3, 4}},
		{[]int{-5, -2, 5, 2, 4, 7, 1, 8, 0, -8}, []int{-5, 5, -2, 2, -8, 4, 7, 1, 8, 0}},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t22(cases[i].test))
	}

}

// Local Variables:
// compile-command: "go test -run Test_t22 -v t22_test.go"
// End:
