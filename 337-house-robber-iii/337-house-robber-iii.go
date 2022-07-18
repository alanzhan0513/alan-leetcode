/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    var dfs func(root *TreeNode) [2]int
	dfs = func(root *TreeNode) [2]int { //0选 1不选
		if root == nil {
			return [2]int{0, 0}
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		selected := root.Val + l[1] + r[1]
		notSelected := max(l[0], l[1]) + max(r[0], r[1])
		return [2]int{selected, notSelected}
	}
	val := dfs(root)
	return max(val[0], val[1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
