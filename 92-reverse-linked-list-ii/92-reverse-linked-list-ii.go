/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    protect := &ListNode{0, head}
    
    node := protect
    for i := 1; i < left; i++ {
        node = node.Next
    }
    
    pre := node
    curr := pre.Next
    tail := curr
    for i := left; i <= right; i++ {
        next := curr.Next
        curr.Next = pre
        pre = curr
        curr = next
    }
    tail.Next = curr
    node.Next = pre
    return protect.Next
}
