/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var list []int
func inorderTraversal(root *TreeNode) []int {
    list = []int{}
    dfs(root)
    return list
}

func dfs(root *TreeNode) {
    if root == nil {
        return
    }
    
    dfs(root.Left)
    list = append(list, root.Val)
    dfs(root.Right)
}