/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
var queue []*Node
var ans [][]int
func levelOrder(root *Node) [][]int {
    if root == nil {
        return [][]int{}
    }
    queue = []*Node{root}
    ans = [][]int{}
    recusive()
    return ans
}

func recusive() {
    length := len(queue)
    if length == 0 {
        return
    }

    temp := []int{}
    for i := 0; i < length; i++ {
        node := queue[0]
        queue = queue[1:]
        
        temp = append(temp, node.Val)
        
        for j := 0; j < len(node.Children); j++ {
            queue = append(queue, node.Children[j])
        }
    }
    ans = append(ans, temp)
    
    recusive()
}