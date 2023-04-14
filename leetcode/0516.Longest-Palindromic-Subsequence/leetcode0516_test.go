package leetcode

import (
	"fmt"
	"testing"
)

type question0516 struct {
	para0516
	ans0516
}

type para0516 struct {
	s string
}

type ans0516 struct {
	one int
}

func Test_Problem0516(t *testing.T) {
	qs := []question0516{
		{
			para0516{"bbbab"},
			ans0516{4},
		}, {
			para0516{"cbbd"},
			ans0516{2},
		},
	}

	fmt.Printf("-----Leetcode Problem 0516-----\n")

	for _, q := range qs {
		out, in := q.ans0516, q.para0516
		got := longestPalindromeSubseq(in.s)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %d", out.one)
		}
		fmt.Printf("【Output】= %d\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
