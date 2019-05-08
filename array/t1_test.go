package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t1(arr []int, n int) []int {
	n = n % len(arr)
	g := gcd(len(arr), n)
	for i := 0; i < g; i++ {
		tmp := arr[i]
		j := i
		for {
			k := j + n
			if k >= len(arr) {
				k -= len(arr)
			}
			if k == i {
				break
			}

			arr[j] = arr[k]
			j = k
		}
		arr[j] = tmp
	}
	return arr
}
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

var rotationCases = []struct {
	test   []int
	n      int
	expect []int
}{
	{[]int{1, 2, 3, 4, 5, 6}, 4, []int{5, 6, 1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5, 6}, 3, []int{4, 5, 6, 1, 2, 3}},
	{[]int{1, 2, 3, 4, 5, 6}, 7, []int{2, 3, 4, 5, 6, 1}},
}

func Test_t1(t *testing.T) {
	for i := range rotationCases {
		testcase := make([]int, len(rotationCases[i].test))
		copy(testcase, rotationCases[i].test)
		assert.Equal(t, rotationCases[i].expect, t1(testcase, rotationCases[i].n))
	}
}
