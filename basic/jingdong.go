package basic

import "golang-acm/util"

/**
链表反转

1 ->2 ->3 > 4

4 3 2 1

 */

func ReverseList(head *util.ListNode) *util.ListNode {
    if head == nil {
        return nil
    }


    left, right := &util.ListNode{}, head
    temp := right.Next
    for right != nil {
        temp = right.Next
        right.Next = left
        left = right
        right = temp
    }

    return left
}
