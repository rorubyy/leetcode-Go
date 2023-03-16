package leetcode

import (
	"fmt"
	"leetcode-Go/structures"
	"testing"
)

type question0106 struct {
	para0106
	ans0106
}

type para0106 struct {
	inorder   []int
	postorder []int
}

type ans0106 struct {
	one []int
}

func Test_Problem0106(t *testing.T) {
	qs := []question0106{

		{
			para0106{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}},
			ans0106{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
		},
	}

	fmt.Printf("-----Leetcode Problem 106-----\n")

	for _, q := range qs {
		_, in := q.ans0106, q.para0106
		fmt.Printf("【Input】= %v\n", in)
		fmt.Printf("【Output】= %v\n", structures.Tree2ints(buildTree(in.inorder, in.postorder)))
	}
	fmt.Printf("\n\n\n")
}
