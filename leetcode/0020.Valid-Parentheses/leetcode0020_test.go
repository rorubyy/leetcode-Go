package leetcode

import (
	"fmt"
	"testing"
)

type question0020 struct {
	para0020
	ans0020
}

type para0020 struct {
	s string
}

type ans0020 struct {
	output bool
}

func Test_Problem0020(t *testing.T) {
	qs := []question0020{
		{
			para0020{"()"},
			ans0020{true},
		},
		{
			para0020{"]"},
			ans0020{false},
		},
		{
			para0020{"(]"},
			ans0020{false},
		},
	}

	fmt.Printf("-----Leetcode Problem 0020-----\n")

	for _, q := range qs {
		out, in := q.ans0020, q.para0020
		got := isValid(in.s)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.output {
			t.Errorf("【Excepted】 = %t", out.output)
		}
		fmt.Printf("【Output】= %t\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
