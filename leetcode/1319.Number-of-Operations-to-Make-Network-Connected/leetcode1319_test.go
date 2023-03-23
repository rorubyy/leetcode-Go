package leetcode

import (
	"fmt"
	"testing"
)

type question1319 struct {
	para1319
	ans1319
}

type para1319 struct {
	n   int
	one [][]int
}

type ans1319 struct {
	one int
}

func Test_Problem1319(t *testing.T) {
	qs := []question1319{
		{
			para1319{4, [][]int{{0, 1}, {0, 2}, {1, 2}}},
			ans1319{1},
		},
	}

	fmt.Printf("-----Leetcode Problem 1319-----\n")

	for _, q := range qs {
		out, in := q.ans1319, q.para1319
		got := makeConnected(in.n, in.one)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %d", out.one)
		}
		fmt.Printf("【Output】= %d\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
