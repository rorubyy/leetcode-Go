package leetcode

import (
	"fmt"
	"testing"
)

type question1466 struct {
	para1466
	ans1466
}

type para1466 struct {
	n   int
	one [][]int
}

type ans1466 struct {
	one int
}

func Test_Problem1466(t *testing.T) {
	qs := []question1466{
		{
			para1466{6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}},
			ans1466{3},
		},
		{
			para1466{5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}},
			ans1466{2},
		},
	}

	fmt.Printf("-----Leetcode Problem 1466-----\n")

	for _, q := range qs {
		out, in := q.ans1466, q.para1466
		got := minReorder(in.n, in.one)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %d", out.one)
		}
		fmt.Printf("【Output】= %d\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
