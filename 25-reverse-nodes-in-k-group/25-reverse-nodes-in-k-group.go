func reverseKGroup(head *ListNode, k int) *ListNode {
    // 1. grouping and find head and tail of group
    // 2. reverse this group
    // 3. update head and tail

    sentinal := &ListNode{}
    last := sentinal
    for head != nil {
        // 1. grouping and find head and tail of group
        end := findEnd(head, k)
        if end == nil {
            break
        }
        nextGroupHead := end.Next
        
        // 2. reverse this group
        reverseLinkedList(head, nextGroupHead)
        
        // 3. update head and tail
        last.Next = end
        head.Next = nextGroupHead
        
        last = head
        head = nextGroupHead
    }
    return sentinal.Next
}

func findEnd(head *ListNode, k int) *ListNode {
    for head != nil {
        k--
        if k == 0 {
            return head
        }
        head = head.Next
    }
    return nil
}

func reverseLinkedList(head, tail *ListNode) {
    last := head
    head = head.Next
    for head != tail {
        headNext := head.Next
        head.Next = last
        last = head
        head = headNext
    }
}