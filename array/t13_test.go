package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t13(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	arr2 := append(arr, arr...)
	max := 0
	for i := 1; i < len(arr); i++ {
		hamming := 0
		for j := 0; j < len(arr); j++ {
			if arr2[j] != arr2[i+j] {
				hamming++
			}
		}
		if max < hamming {
			max = hamming
		}
	}
	return max
}

func Test_t13(t *testing.T) {
	var cases = []struct {
		test   []int
		expect int
	}{
		{[]int{1, 4, 1}, 2},
		{[]int{2, 4, 8, 0}, 4},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t13(cases[i].test))
	}

}

// Local Variables:
// compile-command: "go test -run Test_t13 -v ./..."
// End:
