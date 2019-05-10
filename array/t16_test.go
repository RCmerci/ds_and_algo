package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t16(arr []int, ranges [][]int, index int) int {
	for i := range ranges {
		buf := arr[ranges[i][1]]
		for j := ranges[i][1]; j > ranges[i][0]; j-- {
			arr[j] = arr[j-1]
		}
		arr[ranges[i][0]] = buf
	}
	return arr[index]
}

func Test_t16(t *testing.T) {
	var cases = []struct {
		test   []int
		ranges [][]int
		index  int
		expect int
	}{
		{[]int{1, 2, 3, 4, 5}, [][]int{{0, 2}, {0, 3}}, 1, 3},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t16(cases[i].test, cases[i].ranges, cases[i].index))
	}

}

// Local Variables:
// compile-command: "go test -run Test_t16 -v ./..."
// End:
