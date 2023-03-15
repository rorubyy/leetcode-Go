package leetcode

import "leetcode-Go/structures"

type TreeNode = structures.TreeNode

func isCompleteTree(root *TreeNode) bool {
	queue, found := []*TreeNode{root}, false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			found = true
		} else {
			if found {
				return false
			}
			queue = append(queue, node.Left, node.Right)
		}
	}
	return true
}
