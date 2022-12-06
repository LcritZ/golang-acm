package leetcode

import "golang-acm/util"

/**
删除单链表重复节点
1,1,2,3,3
 */
func DeleteDuplicates(head *util.ListNode) *util.ListNode {
    if head == nil {
        return nil
    }
    slow, fast := head, head
    for fast != nil {
        if fast.Val != slow.Val {
            slow.Next = fast
            slow = fast
        }
        fast = fast.Next
    }

    //打补丁，确保fast最后的处理
    slow.Next = nil

    return head
}
