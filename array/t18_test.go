package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t18(arr []int) []int {
	for i := 0; i < len(arr); {
		if arr[i] > 0 && arr[i] != i {
			tmp := arr[arr[i]]
			arr[arr[i]] = arr[i]
			arr[i] = tmp
		} else {
			i++
		}
	}
	return arr
}

func Test_t18(t *testing.T) {
	var cases = []struct {
		test   []int
		expect []int
	}{
		{[]int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}, []int{-1, 1, 2, 3, 4, -1, 6, -1, -1, 9}},
		{[]int{19, 7, 0, 3, 18, 15, 12, 6, 1, 8,
			11, 10, 9, 5, 13, 16, 2, 14, 17, 4}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19}},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t18(cases[i].test))
	}

}

// Local Variables:
// compile-command: "go test -run Test_t18 -v ./..."
// End:
