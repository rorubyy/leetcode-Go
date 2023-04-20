package leetcode

import (
	"fmt"
	"leetcode-Go/structures"
	"testing"
)

type question0662 struct {
	para0662
	ans0662
}

type para0662 struct {
	one []int
}

type ans0662 struct {
	one int
}

func Test_Problem0662(t *testing.T) {
	qs := []question0662{
		{
			para0662{[]int{1, 3, 2, 5}},
			ans0662{2},
		},
		{
			para0662{[]int{1, 3, 2, 5, 3, 0, 9}},
			ans0662{4},
		},
	}

	fmt.Printf("-----Leetcode Problem 0662-----\n")

	for _, q := range qs {
		out, in := q.ans0662, q.para0662
		root := structures.Ints2TreeNode(in.one)
		got := widthOfBinaryTree(root)
		fmt.Printf("【Input】= %v\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %d", out.one)
		}
		fmt.Printf("【Output】= %v\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
