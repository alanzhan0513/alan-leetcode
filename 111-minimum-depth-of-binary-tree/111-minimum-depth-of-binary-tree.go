/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    depth := 0
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        depth++
        length := len(queue)
        for i := 0; i < length; i++ {
            node := queue[0]
            if node.Left == nil && node.Right == nil {
                return depth
            }
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            queue = queue[1:]
        }
    }
    
    return depth
}