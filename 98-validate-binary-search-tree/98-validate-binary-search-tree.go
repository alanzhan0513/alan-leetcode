/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return check(root, -(1 << 31), 1 << 31 -1)
}

func check(root *TreeNode, min, max int) bool {
    if root == nil {
        return true
    }
    
    if min > root.Val || root.Val > max {
        return false
    }
    
    return check(root.Left, min, root.Val - 1) && check(root.Right, root.Val + 1, max)
}
