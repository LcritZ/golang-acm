package leetcode

import (
    "golang-acm/util"
)

func InsertionSortList(head *util.ListNode) *util.ListNode {
    if head == nil {
        return head
    }

    dummy := &util.ListNode{Next: head}
    prev, curr, temp := dummy, head, head.Next
    for temp != nil {
        if temp.Val >= curr.Val {
            curr = temp
            temp = temp.Next
        } else {
            prev = dummy
            for prev != curr {
                if prev.Next.Val <= temp.Val {
                    prev = prev.Next
                } else {
                    break
                }
            }
            curr.Next = temp.Next
            temp.Next = prev.Next
            prev.Next = temp
            temp = curr.Next
        }
    }

    return dummy.Next
}
