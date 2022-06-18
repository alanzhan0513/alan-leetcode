/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
    "strings"
    "strconv"
)
type Codec struct {
    temp   []string
    current int
}

func Constructor() Codec {
    return Codec{[]string{},0}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    this.temp = []string{}
    this.dfs(root)
    return strings.Join(this.temp, ",")
}

func (this *Codec) dfs(root *TreeNode) {
    if root == nil {
        this.temp = append(this.temp, "null")
        return
    }
    this.temp = append(this.temp, strconv.Itoa(root.Val))
    this.dfs(root.Left)
    this.dfs(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {   
    this.temp = strings.Split(data, ",")
    this.current = 0
    return this.restore()
}

func (this *Codec) restore() *TreeNode {
    if this.temp[this.current] == "null" {
        this.current++
        return nil
    }
    i, _ := strconv.Atoi(this.temp[this.current])
    root := &TreeNode{i, nil, nil}
    this.current++
    root.Left = this.restore()
    root.Right = this.restore()
    return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */