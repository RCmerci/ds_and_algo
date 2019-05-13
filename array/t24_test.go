package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// origin version
func t24(arr []int) []int {
	l, r := 0, 0
	for r < len(arr) && l < len(arr) {
		for l < len(arr) && arr[l] != 0 {
			l++
		}
		if l == len(arr) {
			return arr
		}
		r = l
		for r < len(arr) && arr[r] == 0 {
			r++
		}
		if r == len(arr) {
			return arr
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	return arr
}

// better version
func t24_better(arr []int) []int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count], arr[i] = arr[i], arr[count]
			count++
		}
	}
	return arr
}

func Test_t24(t *testing.T) {
	var cases = []struct {
		test   []int
		expect []int
	}{
		{[]int{1, 2, 0, 4, 3, 0, 5, 0}, []int{1, 2, 4, 3, 5, 0, 0, 0}},
		{[]int{1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0}, []int{1, 9, 8, 4, 2, 7, 6, 0, 0, 0, 0}},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t24_better(cases[i].test))
	}
}

// Local Variables:
// compile-command: "go test -run Test_t24 -v t24_test.go"
// End:
