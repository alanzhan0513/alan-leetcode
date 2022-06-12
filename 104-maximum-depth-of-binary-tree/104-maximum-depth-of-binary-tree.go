/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var depth int
var ans int
func maxDepth(root *TreeNode) int {
    depth = 1
    ans = 0
    calculateDepth(root)
    return ans
}

func calculateDepth(root *TreeNode) {
    if root == nil {
        return
    }
    
    if ans < depth {
        ans = depth
    }
    
    depth++
    calculateDepth(root.Left)
    calculateDepth(root.Right)
    depth--
}
