package leetcode

import "leetcode-Go/structures"

type TreeNode = structures.TreeNode

func buildTreeHelper(inStr int, inEnd int, postStr int, postEnd int, postorder []int, inorderMap map[int]int) *TreeNode {
	if inStr > inEnd || postStr > postEnd {
		return nil
	}
	curRoot := postorder[postEnd]
	curRootIdx := inorderMap[curRoot]
	root := &TreeNode{Val: curRoot}
	leftSubtreeLen := curRootIdx - inStr
	root.Left = buildTreeHelper(inStr, curRootIdx-1, postStr, postStr+leftSubtreeLen-1, postorder, inorderMap)
	root.Right = buildTreeHelper(curRootIdx+1, inEnd, postStr+leftSubtreeLen, postEnd-1, postorder, inorderMap)

	return root

}
func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inorderMap[inorder[i]] = i
	}
	return buildTreeHelper(0, len(inorder)-1, 0, len(inorder)-1, postorder, inorderMap)
}
