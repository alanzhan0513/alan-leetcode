/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    // 1. group
    // 2. reverse
    // 3. update every last and next group's bound
    protect := &ListNode{0, head}
    last := protect
    for head != nil {
        end := getEndListNode(head, k)
        if end == nil {
            break
        }
        
        nextGroupHead := end.Next
        // 2. reverse
        reverseListNode(head, nextGroupHead)
        
        // 3. update every last and next group's bound
        last.Next = end
        head.Next = nextGroupHead
        
        last = head
        head = nextGroupHead
    }
    
    return protect.Next
}

func getEndListNode(head *ListNode, k int) *ListNode {
    for head != nil {
        k--
        if k == 0 {
            return head
        }
        head = head.Next
    }
    return nil
}

func reverseListNode(head *ListNode, end *ListNode) {
    last := head
    head = head.Next
    for head != end {
        nextHead := head.Next
        head.Next = last
        last = head
        head = nextHead
    }
}