package leetcode

import (
	"fmt"
	"testing"
)

type question0605 struct {
	para0605
	ans0605
}

type para0605 struct {
	one []int
	n   int
}

type ans0605 struct {
	one bool
}

func Test_Problem0605(t *testing.T) {
	qs := []question0605{
		{
			para0605{[]int{1, 0, 0, 0, 0, 1}, 2},
			ans0605{false},
		},
		{
			para0605{[]int{1, 0, 0, 0, 1}, 2},
			ans0605{false},
		},
	}

	fmt.Printf("-----Leetcode Problem 605-----\n")

	for _, q := range qs {
		out, in := q.ans0605, q.para0605
		got := canPlaceFlowers(in.one, in.n)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %t", out.one)
		}
		fmt.Printf("【Output】= %t\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
