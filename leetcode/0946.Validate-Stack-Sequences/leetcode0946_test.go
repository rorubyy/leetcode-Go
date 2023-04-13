package leetcode

import (
	"fmt"
	"testing"
)

type question0946 struct {
	para0946
	ans0946
}

type para0946 struct {
	pushed []int
	popped []int
}

type ans0946 struct {
	one bool
}

func Test_Problem0946(t *testing.T) {
	qs := []question0946{
		{
			para0946{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}},
			ans0946{true},
		}, {
			para0946{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}},
			ans0946{false},
		},
	}

	fmt.Printf("-----Leetcode Problem 0946-----\n")

	for _, q := range qs {
		out, in := q.ans0946, q.para0946
		got := validateStackSequences(in.pushed,in.popped)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %t", out.one)
		}
		fmt.Printf("【Output】= %t\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
