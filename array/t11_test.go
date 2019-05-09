package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t11(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return arr[0]
	}
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if arr[mid] > arr[(mid+1)%len(arr)] {
			return arr[(mid+1)%len(arr)]
		}
		if arr[mid] > arr[l] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	panic("")
}

func Test_t11(t *testing.T) {
	var cases = []struct {
		test   []int
		expect int
	}{
		{[]int{15, 18, 2, 3, 6, 12}, 2},
		{[]int{5, 6, 1, 2, 3, 4}, 1},
		{[]int{1, 2, 3, 4, 0}, 0},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t11(cases[i].test))
	}

}

// Local Variables:
// compile-command: "go test -run Test_t11 -v ./..."
// End:
