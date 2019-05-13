package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t19(arr []int) []int {
	l, r := 0, len(arr)-1
	for l < r {
		tmp := arr[l]
		arr[l] = arr[r]
		arr[r] = tmp
		l++
		r--
	}
	return arr
}

func Test_t19(t *testing.T) {
	var cases = []struct {
		test   []int
		expect []int
	}{
		{[]int{4, 5, 1, 2}, []int{2, 1, 5, 4}},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t19(cases[i].test))
	}

}

// Local Variables:
// compile-command: "go test -run Test_t19 -v ./..."
// End:
