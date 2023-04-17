package leetcode

import (
	"fmt"
	"testing"
)

type question1431 struct {
	para1431
	ans1431
}

type para1431 struct {
	candies      []int
	extraCandies int
}

type ans1431 struct {
	output []bool
}

func Test_Problem1431(t *testing.T) {
	qs := []question1431{
		{
			para1431{[]int{2, 3, 5, 1, 3}, 3},
			ans1431{[]bool{true, true, true, false, true}},
		}, {
			para1431{[]int{4, 2, 1, 1, 2}, 1},
			ans1431{[]bool{true, false, false, false, false}},
		},
	}

	fmt.Printf("-----Leetcode Problem 1431-----\n")

	for _, q := range qs {
		out, in := q.ans1431, q.para1431
		got := kidsWithCandies(in.candies, in.extraCandies)
		fmt.Printf("【Input】= %v\n", in)
		for i, v := range out.output {
			fmt.Printf("got: %t, ans:%t\n", got[i], v)
		}
	}
	fmt.Printf("\n\n\n")
}
