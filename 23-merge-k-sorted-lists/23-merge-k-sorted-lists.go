/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    return mergeLists(lists, 0, len(lists) - 1)
}

func mergeLists(lists []*ListNode, low, high int) *ListNode {
    if low == high {
        return lists[low]
    }
    if low > high {
        return nil
    }
    mid := (low + high) / 2
    return merge(mergeLists(lists, low, mid), mergeLists(lists, mid + 1, high))
}

func merge(list1, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    sentinel := &ListNode{}
    last := sentinel
    for list1 != nil && list2 != nil {
        if list1.Val > list2.Val {
            last.Next = list2
            list2 = list2.Next
        } else {
            last.Next = list1
            list1 = list1.Next
        }
        last = last.Next
    }
    if list1 != nil {
        last.Next = list1
    }
    if list2 != nil {
        last.Next = list2
    }
    return sentinel.Next
}