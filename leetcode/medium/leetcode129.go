package leetcode

import 	"github.com/rorubyy6/leetcode-Go/structures"
func sumNumbers(root *TreeNode) int {
    
}

func dfs(root *TreeNode, sum int, ans int){
	if root==nil{
		return
	}
	sum=sum*10+root.Val
	if root.Left==nil && root.right==nil{
		ans+=sum
		return
	}
	dfs(root.Left,sum,ans)
	dfs(root.Right,sum,ans)
}