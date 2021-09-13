package leetcode

import (
	"container/heap"
	"golang-acm/util"
)

//
//func MergeKLists(lists []*util.ListNode) *util.ListNode {
//	if len(lists) == 0 {
//		return nil
//	}
//	if len(lists) == 1 {
//		return lists[0]
//	}
//	newNode := lists[0]
//	for i := 0; i < len(lists)-1; i++ {
//		newNode = basic.Merge(newNode, lists[i+1])
//		for newNode != nil {
//			fmt.Printf(strconv.Itoa(newNode.Val))
//			newNode = newNode.Next
//		}
//		fmt.Println()
//	}
//
//	return newNode
//}
//
//func MergeKLists2(lists []*util.ListNode) *util.ListNode {
//	h := recursion(lists, 0, len(lists)-1)
//	return h
//}
//
//func recursion(lists []*util.ListNode, left, right int) *util.ListNode {
//	if(left == right){
//		return lists[left]
//	}
//	mid := left+(right-left)/2
//	fmt.Println(mid, "--")
//	leftList := recursion(lists, left, mid)
//	rightList := recursion(lists, mid+1, right)
//	//fmt.Println(leftList.Val,rightList.Val)
//	return basic.Merge(leftList, rightList)
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNodes []*util.ListNode

func (l *ListNodes) Len() int {
	return len(*l)
}

func (l *ListNodes) Less(i, j int) bool {
	return (*l)[i].Val < (*l)[j].Val
}

func (l *ListNodes) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *ListNodes) Pop() interface{} {
	n := len(*l)
	x := (*l)[n-1]
	*l = (*l)[:n-1]
	return x
}

func (l *ListNodes) Push(x interface{}) {
	*l = append(*l, x.(*util.ListNode))
}

func mergeKLists3(lists []*util.ListNode) *util.ListNode {
	listNodes := &ListNodes{}
	heap.Init(listNodes)
	for _, v := range lists {
		if v != nil {
			heap.Push(listNodes, v)
		}
	}
	head := &util.ListNode{}
	idx := head
	for listNodes.Len() > 0 {
		val := heap.Pop(listNodes).(*util.ListNode)
		idx.Next = val
		if val.Next != nil {
			heap.Push(listNodes, val.Next)
		}
		idx = idx.Next
	}
	return head.Next
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//
//type PQ []*ListNode
//
//func (p PQ) Len() int { return len(p) }
//func (p PQ) Swap(i, j int) {
//	p[i], p[j] = p[j], p[i]
//}
//func (p PQ) Less(i, j int) bool {
//	return p[i].Val < p[j].Val
//}
//
//func (p *PQ) Push(x interface{}) {
//	node := x.(*ListNode)
//	*p = append(*p, node)
//}
//
//func (p *PQ) Pop() interface{} {
//	old := *p
//	n := len(old)
//	item := old[n-1]
//	*p = old[0 : n-1]
//	return item
//}

//func mergeKLists4(lists []*ListNode) *ListNode {
//	h := &ListNode {
//		Val: -1,
//		Next: nil,
//	}
//	t := h
//	if len(lists) == 0 {
//		return h.Next
//	}
//
//	pq := make(PQ, 0)
//	for i, _ := range lists {
//		if lists[i] != nil {
//			pq = append(pq, lists[i])
//		}
//	}
//	heap.Init(&pq)
//
//	for len(pq) > 0 {
//		// 获取K个链表的最小头
//		item:= heap.Pop(&pq).(*ListNode)
//
//		// 给新链表加节点
//		t.Next = item
//		t = t.Next
//
//		// 将next节点放入最小堆中
//		if item.Next != nil {
//			heap.Push(&pq, item.Next)
//		}
//	}
//
//	return h.Next
//
//}
