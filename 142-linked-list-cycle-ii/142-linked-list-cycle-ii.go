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

    pointer1 := head
    pointer2 := intersect
    
    for pointer1 != pointer2 {
        pointer1 = pointer1.Next
        pointer2 = pointer2.Next
    }
    
    return pointer1
}

func getIntersect(head *ListNode) *ListNode {
    fast := head

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        head = head.Next
        if head == fast {
            return head
        }
    }
    return nil
}