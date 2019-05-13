package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t21(arr []int) []int {
	l, r := 0, len(arr)-1
	for {
		for l < len(arr) && arr[l] < 0 {
			l++
		}
		for r >= 0 && arr[r] > 0 {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
	}

	neg, pos := 0, l
	for neg < pos && pos < len(arr) {
		if arr[neg] < 0 {
			arr[neg], arr[pos] = arr[pos], arr[neg]
		}

		neg += 2
		pos++
	}

	return arr
}

func Test_t21(t *testing.T) {
	var cases = []struct {
		test   []int
		expect []int
	}{
		{[]int{-1, 2, -3, 4, 5, 6, -7, 8, 9}, []int{4, -7, 5, -1, 6, -3, 2, 8, 9}},
		{[]int{-1, -1, -1, 2, 2}, []int{2, -1, 2, -1, -1}},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t21(cases[i].test))
	}

}

// Local Variables:
// compile-command: "go test -run Test_t21 -v t21_test.go"
// End:
