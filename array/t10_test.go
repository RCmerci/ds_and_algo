package array

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func t10(arr []int, rotates []int) []string {
	r := []string{}
	for _, rotate := range rotates {
		rotate %= len(arr)
		sl := []string{}
		for i := 0; i < len(arr); i++ {
			sl = append(sl, strconv.Itoa(arr[(i+rotate+len(arr))%len(arr)]))

		}
		r = append(r, strings.Join(sl, " "))
	}
	return r
}

func Test_t10(t *testing.T) {
	var cases = []struct {
		test1  []int
		test2  []int
		expect []string
	}{
		{[]int{1, 3, 5, 7, 9}, []int{1, 3, 4, 6}, []string{"3 5 7 9 1", "7 9 1 3 5", "9 1 3 5 7", "3 5 7 9 1"}},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expect, t10(cases[i].test1, cases[i].test2))
	}

}
