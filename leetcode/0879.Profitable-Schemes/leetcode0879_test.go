package leetcode

import (
	"fmt"
	"testing"
)

type question0879 struct {
	para0879
	ans0879
}

type para0879 struct {
	n int
	minProfit int
	group []int
	profit []int
}

type ans0879 struct {
	one int
}

func Test_Problem0879(t *testing.T) {
	qs := []question0879{
		{
			para0879{5,3,[]int{2,2},[]int{2,3}},
			ans0879{2},
		},
		{
			para0879{10,5,[]int{2,3,5},[]int{6,7,8}},
			ans0879{7},
		},
	}

	fmt.Printf("-----Leetcode Problem 0879-----\n")

	for _, q := range qs {
		out, in := q.ans0879, q.para0879
		got := profitableSchemes(in.n,in.minProfit,in.group,in.profit)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %d", out.one)
		}
		fmt.Printf("【Output】= %v\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
