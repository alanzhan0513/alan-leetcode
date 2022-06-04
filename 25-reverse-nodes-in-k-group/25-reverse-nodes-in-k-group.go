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
    // 3. update tail and head's ListNode
    sentinel := &ListNode{0, head}
    last := sentinel
    for head != nil {
        end := getEnd(head, k)
        if end == nil {
            break
        }
        
        // 2. reverse ListNode in group
        nextGroupHead := end.Next
        reverse(head, nextGroupHead)
        
        // 3
        last.Next = end
        head.Next = nextGroupHead
        
        
        last = head
        head = nextGroupHead
    }
    return sentinel.Next
}

func getEnd(head *ListNode, k int) *ListNode {
    for head != nil {
        k--
        if k == 0 {
            return head
        }
        head = head.Next
    }
    return nil
}

func reverse(head *ListNode, end *ListNode) {
    last := head
    head = head.Next
    for head != end {
        nextHead := head.Next
        head.Next = last
        last = head
        head = nextHead
    }
}