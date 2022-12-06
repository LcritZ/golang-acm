package leetcode

import (
    "golang-acm/util"
)

//4,2,1,3,8,7,5,6
func SortList(head *util.ListNode) *util.ListNode {
    return sort148(head, nil)
}

//归并，最小集是2个数，head后面就是tail, 那么返回head，同时head.Next=nil， 因为tail在归并的第二个返回了，也就是实际上还是这2个数据merge
func sort148(head, tail *util.ListNode) *util.ListNode {
    if head == nil {
        return head
    }

    if head.Next == tail {
        head.Next = nil
        return head
    }

    slow, fast := head, head
    for fast != tail {
        slow = slow.Next
        fast = fast.Next
        if fast != tail {
            fast = fast.Next
        }
    }

    mid := slow


    return merge148(sort148(head, mid), sort148(mid, tail))

}

func merge148(h1, h2 *util.ListNode) *util.ListNode {
    if h1 == nil && h2 == nil {
        return h1
    }

    dummy := &util.ListNode{}
    curr := dummy
    for h1 != nil && h2 != nil {
        if h1.Val <= h2.Val {
            curr.Next = h1
            h1 = h1.Next
        } else {
            curr.Next = h2
            h2 = h2.Next
        }
        curr = curr.Next
    }

    if h1 != nil {
        curr.Next = h1
    } else if h2 != nil{
        curr.Next = h2
    }

    return dummy.Next

}


