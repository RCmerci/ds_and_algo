package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// recursive version
func t3(arr []int, n int) []int {
	n %= len(arr)

	if 2*n == len(arr) {
		return append(append([]int{}, arr[n:]...), arr[:n]...)
	}
	if n > len(arr)/2 {
		AL := append([]int{}, arr[:len(arr)-n]...)
		AR := append([]int{}, arr[len(arr)-n:n]...)
		B := append([]int{}, arr[n:]...)
		return append(B, t3(append(AR, AL...), len(AR))...)
	}
	A := append([]int{}, arr[:n]...)
	BL := append([]int{}, arr[n:len(arr)-n]...)
	BR := append([]int{}, arr[len(arr)-n:]...)
	return append(t3(append(BR, BL...), len(BR)), A...)
}

// iter version
func t3_iter(arr []int, n int) []int {
	n %= len(arr)
	i, j := 0, len(arr)
	for i < j {
		if 2*n == j-i {
			swap(arr, i, i+n, n)
			return arr
		}
		if n > (j-i)/2 { // AL, AR, B
			swap(arr, i, i+n, j-i-n)
			oldn := n
			n = 2*n - j + i
			i += j - i - oldn

		} else { // A, BL, BR
			swap(arr, i, j-n, n)
			j -= n
		}
	}

	return arr
}

// swap slice[i:i+len] with slice[j:j+len]
func swap(slice []int, i, j, len int) {
	for k := 0; k < len; k++ {
		tmp := slice[i+k]
		slice[i+k] = slice[j+k]
		slice[j+k] = tmp
	}
}

func Test_t3(t *testing.T) {
	for i := range rotationCases {
		testcase := make([]int, len(rotationCases[i].test))
		copy(testcase, rotationCases[i].test)
		assert.Equal(t, rotationCases[i].expect, t3(testcase, rotationCases[i].n))
	}
	for i := range rotationCases {
		testcase := make([]int, len(rotationCases[i].test))
		copy(testcase, rotationCases[i].test)
		assert.Equal(t, rotationCases[i].expect, t3_iter(testcase, rotationCases[i].n))
	}
}
