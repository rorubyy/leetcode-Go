package leetcode

import (
	"fmt"
	"testing"
)

type question0071 struct {
	para0071
	ans0071
}

type para0071 struct {
	path string
}

type ans0071 struct {
	result string
}

func Test_Problem0071(t *testing.T) {
	qs := []question0071{
		{
			para0071{"/home/"},
			ans0071{"/home"},
		},
		{
			para0071{"/../"},
			ans0071{"/"},
		},
		{
			para0071{"/home//foo/"},
			ans0071{"/home/foo"},
		},
	}

	fmt.Printf("-----Leetcode Problem 0071-----\n")

	for _, q := range qs {
		out, in := q.ans0071, q.para0071
		got := simplifyPath(in.path)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.result {
			t.Errorf("【Excepted】 = %v", out.result)
		}
		fmt.Printf("【Output】= %v\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
