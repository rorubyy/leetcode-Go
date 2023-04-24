package leetcode

import (
	"fmt"
	"testing"
)

type question1046 struct {
	para1046
	ans1046
}

type para1046 struct {
	one []int
}

type ans1046 struct {
	one int
}

func Test_Problem1046(t *testing.T) {
	qs := []question1046{
		{
			para1046{[]int{2, 7, 4, 1, 8, 1}},
			ans1046{1},
		},
		{
			para1046{[]int{1}},
			ans1046{1},
		},
		{
			para1046{[]int{2, 2}},
			ans1046{0},
		},
	}

	fmt.Printf("-----Leetcode Problem 958-----\n")

	for _, q := range qs {
		out, in := q.ans1046, q.para1046
		got := lastStoneWeight(in.one)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %d", out.one)
		}
		fmt.Printf("【Output】= %d\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
