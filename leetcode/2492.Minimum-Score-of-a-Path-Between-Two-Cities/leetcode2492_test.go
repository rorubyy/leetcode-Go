package leetcode

import (
	"fmt"
	"testing"
)

type question2492 struct {
	para2492
	ans2492
}

type para2492 struct {
	n   int
	one [][]int
}

type ans2492 struct {
	one int
}

func Test_Problem2492(t *testing.T) {
	qs := []question2492{
		{
			para2492{4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}},
			ans2492{5},
		},
		{
			para2492{4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}},
			ans2492{2},
		},
	}

	fmt.Printf("-----Leetcode Problem 2492-----\n")

	for _, q := range qs {
		out, in := q.ans2492, q.para2492
		got := minScore(in.n, in.one)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %d", out.one)
		}
		fmt.Printf("【Output】= %d\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}

