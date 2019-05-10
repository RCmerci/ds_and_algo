package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func t14(arr []int, queries [][]int) []int {
	if len(arr) == 0 {
		panic("")
	}
	arr2 := make([]int, len(arr))
	arr2[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		arr2[i] = arr2[i-1] + arr[i]
	}

	L := 0
	res := []int{}
	for i := 0; i < len(queries); i++ {
		switch queries[i][0] {
		case 1:
			L = (L + queries[i][1]) % len(arr)
		case 2:
			L = (L - queries[i][1] + len(arr)) % len(arr)
		case 3:
			l := queries[i][1]
			r := queries[i][2]

			res = append(res, (arr2[len(arr)-1]+arr2[((len(arr)-L)%len(arr)+r)%len(arr)]-arr2[((len(arr)-L-1)%len(arr)+l)%len(arr)])%arr2[len(arr)-1])
		}
	}
	return res
}

func Test_t14(t *testing.T) {
	var cases = []struct {
		test    []int
		queries [][]int
		expect  []int
	}{
		{[]int{1, 2, 3, 4, 5}, [][]int{{1, 3}, {3, 0, 2}, {2, 1}, {3, 1, 4}}, []int{12, 11}},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t14(cases[i].test, cases[i].queries))
	}

}

// Local Variables:
// compile-command: "go test -run Test_t14 -v ./..."
// End:
