package leetcode

import (
    "golang-acm/util"
)

/**
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

1,3,2,4,5
  tail tmp tmp =
*/

func ReverseBetween(head *util.ListNode, left int, right int) *util.ListNode {
    if head == nil {
        return nil
    }

    dummy := &util.ListNode{Next: head}
    h1, h2, tail := head, dummy, dummy
    for i := 0; i < right; i++ {
        h1 = h1.Next
        tail = tail.Next
    }
    for i := 0; i < left-1; i++ {
        h2 = h2.Next
    }
    tail.Next = nil
    tmp := h2.Next
    h2.Next = ReverseList(h2.Next)
    tmp.Next = h1
    for tmp != nil {
        tmp = tmp.Next
    }
    return dummy.Next
}

func ReverseBetween2(head *util.ListNode, left int, right int) *util.ListNode {
    if head == nil {
        return nil
    }

    dummy := &util.ListNode{Next: head}
    pre, curr := dummy, head

    for i := 0; i < left-1; i++ {
        pre = pre.Next
        curr = curr.Next
    }

    for i := 0;i < right-left; i++ {
        //指针的指向是关键，如何保存，如何改变
        tail := curr.Next
        curr.Next = tail.Next
        tail.Next = pre.Next
        pre.Next = tail
    }
    return dummy.Next

}
