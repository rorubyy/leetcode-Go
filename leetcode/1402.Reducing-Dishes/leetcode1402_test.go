package leetcode

import (
	"fmt"
	"testing"
)

type question1402 struct {
	para1402
	ans1402
}

type para1402 struct {
	one []int
}

type ans1402 struct {
	one int
}

func Test_Problem1402(t *testing.T) {
	qs := []question1402{
		{
			para1402{[]int{-1, -4, -5}},
			ans1402{0},
		},
		{
			para1402{[]int{-1, -8, 0, 5, -9}},
			ans1402{14},
		},
		{
			para1402{[]int{4, 3, 2}},
			ans1402{20},
		},
	}

	fmt.Printf("-----Leetcode Problem 1402-----\n")

	for _, q := range qs {
		out, in := q.ans1402, q.para1402
		got := maxSatisfaction(in.one)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %d", out.one)
		}
		fmt.Printf("【Output】= %d\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
