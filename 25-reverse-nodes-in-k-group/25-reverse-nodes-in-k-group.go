/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    // 1. grouping and find head and tail of group
    // 2. reverse in the group
    // 3. update head and tail link

    sentinal := &ListNode{}
    last := sentinal
    for head != nil {
        end := findEnd(head, k)
        if end == nil {
            break
        }
        nextGroupHead := end.Next
        
        reverseList(head, nextGroupHead)
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

func reverseList(head, tail *ListNode) {
    last := head
    head = head.Next
    for head != tail {
        headNext := head.Next
        head.Next = last
        last = head
        head = headNext
    }
}