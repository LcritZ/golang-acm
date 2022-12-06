package leetcode

/**
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。

思路：根据总和，算出每一组和 = nums/k
然后回溯出组合，有一组就true

 */

//
//var enable bool
//var groups int
//var numMap map[int]int
//
//func canPartitionKSubsets(nums []int, k int) bool {
//	if len(nums) == 0 || k == 0 {
//		return false
//	}
//	if k == 1 {
//		return true
//	}
//
//	sum, subSum := 0, 0
//	//存入map,统计
//	for i := 0; i<len(nums); i++ {
//		sum += nums[i]
//		numMap[nums[i]]++
//	}
//	if sum%k != 0 {
//		return false
//	} else {
//		subSum = sum/k
//	}
//
//	groups = k
//	enable = false
//
//
//	sort.Ints(nums)
//
//
//
//
//
//}
//
//func trace698(nums []int, sum, groupNum int)   {
//	if len(nums) == 0 && groupNum == groups {
//		enable = true
//		return
//	}
//
//	for i := 0; i < len(nums); i++ {
//		cur := nums[i]
//		if cur > sum {
//			enable = false
//			return
//		} else if cur == sum {
//			groupNum++
//			if i < len(nums)-1 {
//				nums = append(nums[0:i], nums[i+1:]...)
//			} else {
//				nums = []int{}
//			}
//			trace698(nums, sum, groupNum)
//			groupNum--
//			nums = append(nums[:i],append([]int{cur},nums[i:]...)...)
//		} else {
//			//通过map找到下一个
//		}
//	}
//
//}
