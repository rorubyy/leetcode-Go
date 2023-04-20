package leetcode

import "leetcode-Go/structures"

type TreeNode = structures.TreeNode

func maxVal(x uint32, y uint32) uint32 {
	if x > y {
		return x
	} else {
		return y
	}
}

func widthOfBinaryTree(root *TreeNode) int {
	queue, val := []*TreeNode{root}, []uint32{1}
	var maxWidth uint32
	for len(queue) > 0 {
		strIdx := val[0]
		endIdx := val[len(val)-1]
		levelSize := len(queue)
		maxWidth = maxVal(endIdx-strIdx+1, maxWidth)
		for levelSize > 0 {
			levelSize--
			cur := queue[0]
			curIdx := val[0]
			if cur.Left != nil {
				val = append(val, curIdx*2)
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				val = append(val, curIdx*2+1)
				queue = append(queue, cur.Right)

			}
			val = val[1:]
			queue = queue[1:]
		}
	}
	return int(maxWidth)

}
