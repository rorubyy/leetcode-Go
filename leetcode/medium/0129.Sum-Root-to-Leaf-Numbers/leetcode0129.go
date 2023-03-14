package leetcode

import (
	"leetcode-Go/structures"
)

type TreeNode = structures.TreeNode

func sumNumbers(root *TreeNode) int {
	ans := 0
	dfs(root, 0, &ans)
	return ans
}

func dfs(root *TreeNode, sum int, ans *int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*ans += sum
		return
	}
	dfs(root.Left, sum, ans)
	dfs(root.Right, sum, ans)
}
