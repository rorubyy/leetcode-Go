package leetcode

import (
	"fmt"
	"testing"
)

type question0649 struct {
	para0649
	ans0649
}

type para0649 struct {
	one string
}

type ans0649 struct {
	one string
}

func Test_Problem0649(t *testing.T) {
	qs := []question0649{
		{
			para0649{"RDD"},
			ans0649{"Dire"},
		},
		{
			para0649{"RD"},
			ans0649{"Radiant"},
		},
	}

	fmt.Printf("-----Leetcode Problem 0649-----\n")

	for _, q := range qs {
		out, in := q.ans0649, q.para0649
		got := predictPartyVictory(in.one)
		fmt.Printf("【Input】= %s\n", in)
		if got != out.one {
			t.Errorf("【Excepted】 = %s", out.one)
		}
		fmt.Printf("【Output】= %s\n", got)
		fmt.Printf("--------------------\n")
	}
	fmt.Printf("\n\n\n")
}
