package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t23(arr []int) []int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}
	for i := count; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr
}

func Test_t23(t *testing.T) {
	var cases = []struct {
		test   []int
		expect []int
	}{
		{[]int{1, 2, 0, 4, 3, 0, 5, 0}, []int{1, 2, 4, 3, 5, 0, 0, 0}},
		{[]int{1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0}, []int{1, 9, 8, 4, 2, 7, 6, 0, 0, 0, 0}},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t23(cases[i].test))
	}

}

// Local Variables:
// compile-command: "go test -run Test_t23 -v t23_test.go"
// End:
