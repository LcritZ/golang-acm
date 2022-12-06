package leetcode

import (
    "fmt"
    "golang-acm/util"
)

func DetectCycle(head *util.ListNode) *util.ListNode {
    if head == nil {
        return head
    }

    dummy := &util.ListNode{Next: head}
    slow, fast := dummy, dummy
    var hasCycle bool
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            hasCycle = true
            break
        }
    }

    if !hasCycle {
        return nil
    }
    fmt.Println(slow.Val)
    temp := dummy
    for temp != slow {
        temp = temp.Next
        slow = slow.Next
    }

    return temp
}