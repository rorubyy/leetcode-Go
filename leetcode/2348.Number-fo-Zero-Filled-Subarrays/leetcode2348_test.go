package leetcode

import (
	"fmt"
	"testing"
)

type question2348 struct {
	para2348
	ans2348
}

type para2348 struct {
	one []int
}

type ans2348 struct {
	one int64
}

func Test_Problem2348(t *testing.T) {
	qs := []question2348{
		{
			para2348{[]int{1, 3, 0, 0, 2, 0, 0, 4}},
			ans2348{6},
		},
		{
			para2348{[]int{0, 0, 0, 2, 0, 0}},
			ans2348{9},
		},
	}

	fmt.Printf("-----Leetcode Problem 2348-----\n")

	for _, q := range qs {
		out, in := q.ans2348, q.para2348
		got := zeroFilledSubarray(in.one)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %d", out.one)
		}
		fmt.Printf("【Output】= %d\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
