package leetcode

import "golang-acm/util"

/**
思路：链表存储到map，然后看从哪个节点开始数字一样，每个链表走到数字相同前面，开始判断节点的地址是不是同一个

解法：
 */
func GetIntersectionNode(headA, headB *util.ListNode) *util.ListNode {

    if headA == nil || headB == nil {
        return nil
    }
    h1, h2 := headA, headB
    for h1 != h2 {
        if h1 == nil {
            h1 = headB
        } else {
            h1 = h1.Next
        }
        if h2 == nil {
            h2 = headA
        } else {
            h2 = h2.Next
        }
    }
    return h1
}
