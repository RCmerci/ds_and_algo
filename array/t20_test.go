package array

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func t20(arr []int) []int {
	sort.Ints(arr)
	r := make([]int, len(arr))
	for i := 0; 2*i < len(arr); i++ {
		idx := len(arr) - 1 - 2*i
		if idx%2 == 0 { // odd
			r[idx] = arr[i]
			if idx-1 >= 0 {
				r[idx-1] = arr[len(arr)-1-i]
			}
		} else { // even
			r[idx] = arr[len(arr)-1-i]
			if idx-1 >= 0 {
				r[idx-1] = arr[i]
			}
		}
	}
	return r

}

func Test_t20(t *testing.T) {
	var cases = []struct {
		test   []int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{4, 5, 3, 6, 2, 7, 1}},
		{[]int{1, 2, 1, 4, 5, 6, 8, 8}, []int{4, 5, 2, 6, 1, 8, 1, 8}},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t20(cases[i].test))
	}

}

// Local Variables:
// compile-command: "go test -run Test_t20 -v ./..."
// End:
