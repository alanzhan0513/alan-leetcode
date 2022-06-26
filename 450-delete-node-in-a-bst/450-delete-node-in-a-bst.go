/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return root
    }
    
    if key == root.Val {
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }
        
        next := root.Right
        for next.Left != nil {
            next = next.Left
        }
        root.Right = deleteNode(root.Right, next.Val)
        root.Val = next.Val
        return root
    }
    
    if root.Val < key {
        root.Right = deleteNode(root.Right, key)
    } else {
        root.Left = deleteNode(root.Left, key)
    }
    return root
}