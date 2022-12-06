package leetcode

import "golang-acm/util"

//递归处理，之一当前是否满足K个，满足就对前K个进行翻转，然后递归拿到后面的链表
func reverseKGroup(head *util.ListNode, k int) *util.ListNode {
    if head == nil {
        return nil
    }

    dummy := &util.ListNode{Next: head}

    temp := dummy
    i := 0
    for ; i < k && temp!= nil; i++ {
        temp = temp.Next
    }
    if temp == nil {
        return dummy.Next
    }
    //翻转
    nextListHead := temp.Next
    temp.Next = nil

    newTail := head
    slow, fast := dummy, head
    for fast != nil {
        tail := fast.Next
        fast.Next = slow
        slow = fast
        fast = tail
    }
    dummy.Next = slow
    newTail.Next = reverseKGroup(nextListHead, k)

    return dummy.Next
}

