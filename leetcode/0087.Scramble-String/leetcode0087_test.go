package leetcode

import (
	"fmt"
	"testing"
)

type question0087 struct {
	para0087
	ans0087
}

type para0087 struct {
	s1 string
	s2 string
}

type ans0087 struct {
	output bool
}

func Test_Problem0087(t *testing.T) {
	qs := []question0087{
		{
			para0087{"great", "rgeat"},
			ans0087{true},
		},
		{
			para0087{"abcde", "caebd"},
			ans0087{false},
		},
		{
			para0087{"a", "a"},
			ans0087{true},
		},
	}

	fmt.Printf("-----Leetcode Problem 0087-----\n")

	for _, q := range qs {
		out, in := q.ans0087, q.para0087
		got := isScramble(in.s1, in.s2)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.output {
			t.Errorf("【Excepted】 = %t", out.output)
		}
		fmt.Printf("【Output】= %t\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
