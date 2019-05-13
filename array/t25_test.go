package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t25(arr []int, k int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] <= k {
			count++
		}
	}

	set := map[int]int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] <= k {
			for j := range set {
				if i-j < count {
					set[j]++
				}
			}
			set[i] = 1
		}
	}

	max := 0
	for i := range set {
		if max < set[i] {
			max = set[i]
		}
	}
	return count - max
}

func t25_slide_window(arr []int, k int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] <= k {
			count++
		}
	}
	bad := 0
	for i := 0; i < count; i++ {
		if arr[i] > k {
			bad++
		}
	}
	ans := bad
	j := count
	for i := 0; i < len(arr); i++ {
		if j == len(arr) {
			break
		}
		if arr[i] > k {
			bad--
		}
		if arr[j] > k {
			bad++
		}
		if ans > bad {
			ans = bad
		}
		j++
	}
	return ans
}

func Test_t25(t *testing.T) {
	var cases = []struct {
		test   []int
		k      int
		expect int
	}{
		{[]int{2, 1, 5, 6, 3}, 3, 1},
		{[]int{2, 7, 9, 5, 8, 7, 4}, 5, 2},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t25(cases[i].test, cases[i].k))
		assert.Equal(t, cases[i].expect, t25_slide_window(cases[i].test, cases[i].k))
	}
}

// Local Variables:
// compile-command: "go test -run Test_t25 -v t25_test.go"
// End:
