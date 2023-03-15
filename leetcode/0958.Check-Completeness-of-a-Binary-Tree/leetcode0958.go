package leetcode

import "leetcode-Go/structures"

type TreeNode = structures.TreeNode

func isCompleteTree(root *TreeNode) bool {
	queue, found := []*TreeNode{root}, false //initialize a queue save tree node
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] //pop the first node from queue
		if node == nil {
			found = true
		} else {
			if found {
				return false
			} // if find a node after a nil -> not a completed tree
			queue = append(queue, node.Left, node.Right)
		}
	}
	return true
}
