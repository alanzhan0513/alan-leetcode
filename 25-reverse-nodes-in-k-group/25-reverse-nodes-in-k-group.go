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
    // 3. update pre next and last next
    var protect = &ListNode {
        Val: 0,
        Next: head,
    }
    last := protect
    for head != nil {
        // 1, group
        end := getEnd(head, k)
        if end == nil {
            break;
        }
        nextGroupHead := end.Next
        
        // 2. reverse
        reverseList(head, nextGroupHead)
        
        // 3. update pre next and last next
        last.Next = end
        head.Next = nextGroupHead
        last = head
        head = nextGroupHead
    }
    return protect.Next
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

func reverseList(head *ListNode, end *ListNode) {
    last := head
    head = head.Next
    for head != end {
        nextHead := head.Next
        head.Next = last
        last = head
        head = nextHead
    }
}