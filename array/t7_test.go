package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t7(arr []int) int {
	arrSum := 0
	for i := range arr {
		arrSum += arr[i]
	}
	curr := 0
	for i := range arr {
		curr += i * arr[i]
	}
	max := curr
	for i := 1; i < len(arr); i++ {
		curr += arrSum - len(arr)*arr[len(arr)-i]
		if curr > max {
			max = curr
		}
	}
	return max
}

func Test_t7(t *testing.T) {
	var cases = []struct {
		test   []int
		expect int
	}{
		{[]int{1, 20, 2, 10}, 72},
		{[]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 330},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t7(cases[i].test))
	}

}
