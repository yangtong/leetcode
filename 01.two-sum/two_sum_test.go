package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	nums   []int
	target int
}

type ans struct {
	one []int
}

func TestTwoSum(t *testing.T) {
	c := assert.New(t)

	tabletests := []question{
		{
			para{[]int{3, 2, 4}, 6},
			ans{[]int{1, 2}},
		},

		{
			para{[]int{3, 2, 4}, 5},
			ans{[]int{0, 1}},
		},

		{
			para{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans{[]int{1, 3}},
		},

		{
			para{[]int{0, 1}, 1},
			ans{[]int{0, 1}},
		},

		{
			para{[]int{0, 3}, 5},
			ans{[]int{}},
		},
	}
	for _, tt := range tabletests {
		ans := TwoSum(tt.para.nums, tt.para.target)
		c.Equal(ans, tt.ans.one)

	}

}
