/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    intersect := getIntersect(head)
    if intersect == nil {
        return nil
    }
    
    for head != intersect {
        head = head.Next
        intersect = intersect.Next
    }
    
    return intersect
}

func getIntersect(head *ListNode) *ListNode {
    fast := head
    slow := head
    
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        
        if fast == slow {
            return fast
        }
    }
    return nil
}