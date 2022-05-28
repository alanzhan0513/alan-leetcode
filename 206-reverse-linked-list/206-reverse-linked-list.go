/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var last *ListNode
    
    for head != nil {
        nextHead := head.Next
        head.Next = last
        last = head
        head = nextHead
    }
    return last
}