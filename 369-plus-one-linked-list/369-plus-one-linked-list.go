/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
    sentinel := &ListNode{0, head}
    notNine := sentinel
    
    for head != nil {
        if head.Val != 9 {
            notNine = head
        }
        head = head.Next
    }

    notNine.Val++
    notNine = notNine.Next
    
    for notNine != nil {
        notNine.Val = 0
        notNine = notNine.Next
    }
    
    
    if sentinel.Val != 0 {
        return sentinel
    }
    return sentinel.Next
}