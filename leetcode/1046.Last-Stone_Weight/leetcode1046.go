package leetcode

import (
	"container/heap"
	"leetcode-Go/structures"
)

func lastStoneWeight(stones []int) int {
	h := new(structures.MaxHp)
	heap.Init(h)
	for _, val := range stones {
		heap.Push(h, val)
	}
	for h.Len() > 1 {
		first := (*h)[0]
		heap.Pop(h)
		second := (*h)[0]
		heap.Pop(h)
		if first != second {
			heap.Push(h, first-second)
		}
	}
	if h.Len() > 0 {
		return (*h)[0]
	} else {
		return 0
	}
}
