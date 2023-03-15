package leetcode

import (
	"fmt"
	"leetcode-Go/structures"
	"testing"
)

type question0129 struct {
	para0129
	ans0129
}

type para0129 struct {
	one []int
}

type ans0129 struct {
	one int
}

func Test_Problem0129(t *testing.T) {
	qs := []question0129{
		{
			para0129{[]int{}},
			ans0129{0},
		},
		{
			para0129{[]int{1, 2, 3}},
			ans0129{25},
		},
		{
			para0129{[]int{4, 9, 0, 5, 1}},
			ans0129{1026},
		},
	}

	fmt.Printf("-----Leetcode Problem 129-----\n")

	for _, q := range qs {
		out, in := q.ans0129, q.para0129
		root := structures.Ints2TreeNode(in.one)
		got := sumNumbers(root)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %d", out.one)
		}
		fmt.Printf("【Output】= %v\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
