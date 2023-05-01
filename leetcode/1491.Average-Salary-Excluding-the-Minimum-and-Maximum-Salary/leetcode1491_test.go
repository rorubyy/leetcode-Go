package leetcode

import (
	"fmt"
	"testing"
)

type question1491 struct {
	para1491
	ans1491
}

type para1491 struct {
	one []int
}

type ans1491 struct {
	one float64
}

func Test_Problem1491(t *testing.T) {
	qs := []question1491{
		{
			para1491{[]int{4000,3000,1000,2000}},
			ans1491{2500.00000},
		},
		{
			para1491{[]int{1000,2000,3000}},
			ans1491{2000.00000},
		},
	}

	fmt.Printf("-----Leetcode Problem 1491-----\n")

	for _, q := range qs {
		out, in := q.ans1491, q.para1491
		got := average(in.one)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %v", out.one)
		}
		fmt.Printf("【Output】= %v\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
