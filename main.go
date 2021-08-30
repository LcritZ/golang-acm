package main

import (
	"fmt"
	"golang-acm/common"
	"golang-acm/leetcode"
)

func main() {

	// 342 465
	l1 := &common.ListNode{
		Val: 2,
	}
	l1.Next = &common.ListNode{
		Val: 4,
	}
	l1.Next.Next = &common.ListNode{
		Val: 3,
	}

	l2 := &common.ListNode{
		Val: 5,
	}
	l2.Next = &common.ListNode{
		Val: 6,
	}
	l2.Next.Next = &common.ListNode{
		Val: 4,
	}
	fmt.Println((l1.Val+l2.Val)%10)
	//leetcode.AddTwoNumbers2(l1, l2)
	newHead := leetcode.AddTwoNumbers2(l1, l2)
	for newHead != nil {
		fmt.Println(newHead.Val, "--")
		newHead = newHead.Next
	}
	//fmt.Println(newHead.Val)
	//fmt.Println(newHead.Next.Val)
	//fmt.Println(newHead.Next.Next.Val)
	//fmt.Println(newHead.Next.Next.Next.Val)

	//fmt.Println("---")
	//common.Max(1,2)
	//
	//nums := []int{1,3,5,7,8}
	//fmt.Println(leetcode.TwoSum(nums, 9))
}
