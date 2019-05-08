package array

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func t2(arr []int, n int) []int {
	n %= len(arr)
	A, B := make([]int, n), make([]int, len(arr)-n)
	copy(A, arr[:n])
	copy(B, arr[n:])
	sort.Reverse(sort.IntSlice(A))
	sort.Reverse(sort.IntSlice(B))
	C := append(B, A...)
	sort.Reverse(sort.IntSlice(C))
	return C
}

func Test_t2(t *testing.T) {
	for i := range rotationCases {
		testcase := make([]int, len(rotationCases[i].test))
		copy(testcase, rotationCases[i].test)
		assert.Equal(t, rotationCases[i].expect, t2(testcase, rotationCases[i].n))
	}
}
