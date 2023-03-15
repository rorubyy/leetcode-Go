package leetcode

import (
	"fmt"
	"leetcode-Go/structures"
	"testing"
)

type question0958 struct {
	para0958
	ans0958
}

type para0958 struct {
	one []int
}

type ans0958 struct {
	one bool
}

func Test_Problem0958(t *testing.T) {
	qs := []question0958{
		{
			para0958{[]int{1,2,3,4,5,6}},
			ans0958{true},
		},
		{
			para0958{[]int{1,2,3,4,5,structures.NULL,7}},
			ans0958{false},
		},
	}

	fmt.Printf("-----Leetcode Problem 958-----\n")

	for _, q := range qs {
		out, in := q.ans0958, q.para0958
		root := structures.Ints2TreeNode(in.one)
		got := isCompleteTree(root)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %t", out.one)
		}
		fmt.Printf("【Output】= %t\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
