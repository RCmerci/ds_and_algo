package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t4(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	tmp := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		arr[i+1] = arr[i]
	}
	arr[0] = tmp
	return arr
}

func Test_t4(t *testing.T) {
	var cases = []struct {
		test   []int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 1, 2, 3, 4, 5}},
	}

	for i := range cases {
		testcase := make([]int, len(cases[i].test))
		copy(testcase, cases[i].test)
		assert.Equal(t, cases[i].expect, t4(testcase))
	}

}
