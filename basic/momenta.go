package basic

import (
    "fmt"
    "golang-acm/util"
    "math"
)


// 1.有一个数组，从小到大的，在某个位置发生了翻转
// [1,2,3,4,5]
// [4,5,1,2,3]
// [4,5,6,7,1,2,3]

func getPostion(nums []int) int {
    if len(nums) == 0 {
        return -1
    }

    left, right := 0, len(nums)-1
    min := math.MaxInt32
    pos := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < nums[left] {
            right = mid
        } else {
            left = mid+1
        }
        if nums[mid] < min {
            pos = mid
            min = nums[mid]
        }
    }
    return pos
}

func main() {
    // nums := []int{4,5,1,2,3}
    // nums := []int{4,5,6,1,2,3}
    nums := []int{4,5,6,7,1,2,3}
    fmt.Println(getPostion(nums))
}



// 告诉我在哪里发生了翻转

// 2.二叉树的宽度遍历（层次遍历）

func travseTree(root *util.TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }

    nodeArr := []*util.TreeNode{root}
    for len(nodeArr) > 0 {
        node := nodeArr[0]
        nodeArr = nodeArr[1:]
        res = append(res, node.Val)
        if node.Left != nil {
            nodeArr = append(nodeArr, node.Left)
        }
        if node.Right != nil {
            nodeArr = append(nodeArr, node.Right)
        }
    }

    return res
}






// 有两个从小到大的链表，合并成一个从大到小的链表

// 1 - 3 - 5
// 2 - 4 - 6
// ... 2 - 1




// type ListNode struct {
//   val int
//   next *ListNode
// }


// func merge(l1, l2 *ListNode) *ListNode {
//   if l1 == nil && l2 == nil {
//     return nil
//   }

//   if l1 == nil {
//     return reverse(l2)
//   }
//   if l2 == nil {
//     return reverse(l1)
//   }

//   var prev *ListNode
//   var tail *ListNode
//   var h    *ListNode

//   for l1 != nil && l2 != nil {
//     if l1.val < l2.val {
//       tail = l1
//       if l1.next != nil && l1.next.val < l2.val {
//         prev, h := mergeNode(l1, l1.next, h)
//       } else if l1.next != nil && l1.next.val > l2.val {
//         prev, h := mergeNode(l1, l2, h)
//       }



//     } else {
//       tail = l2
//     }

//     if l1.


//     prevnext = prev.next
//     tailnext = tail.next
//     prev.next = tail
//     tail.next =


//   }


// }

// func mergeNode(n1, n2, h *ListNode) (*ListNode, *ListNode) {
//   var next *ListNode
//   var temp *ListNode
//   if n1.val < n2.val {
//     next = n1.next
//     n1.next = h
//     temp = n1
//   } else {
//     next = n2.next
//     n2.next = h
//     temp = n2
//   }

//   return temp, next

// }

// func reverse(l *ListNode) *ListNode {
//   //todo
// }
