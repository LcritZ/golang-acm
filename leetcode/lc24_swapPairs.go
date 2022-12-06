package leetcode

import "golang-acm/util"

func swapPairs(head *util.ListNode) *util.ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    dummy := &util.ListNode{}
    newHead := head.Next.Next

    temp := head.Next
    dummy.Next = temp
    temp.Next = head
    head.Next = swapPairs(newHead)
    return dummy.Next
}

