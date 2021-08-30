package acm

import "golang-acm/common"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param head ListNode类
 * @return int整型一维数组
 */
func PrintListFromTailToHead( head *common.ListNode) []int {
	// write code here
	var resArray = []int{}
	if head != nil {
		resArray = append(resArray, PrintListFromTailToHead(head.Next)...)
		//fmt.Printf("%d", head.Val)
		resArray = append(resArray, head.Val)
	}
	return resArray
}
