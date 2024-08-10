package main

import (
	"fmt"
	"golang-acm/leetcode"
	"golang-acm/util"
)

func main() {

	//res := leetcode.GetModifiedArray(5, [][]int{{1,3,2}, {2,4,3}, {0,2,-2}})
	//fmt.Println(res)

	//n := 4
	//nums := [][]int{{1,2}, {1,3}, {2,4}}
	//res := leetcode.PossibleBipartition(n, nums)
	//fmt.Println(res)

	//nums := []int{4,3,2,6}
	//res := leetcode.MaxRotateFunction(nums)
	//fmt.Println(res)

	//fmt.Println(leetcode.LargestPalindrome(2))
	//998001
	//997997
	//n := 994499
	//fmt.Println(n)
	//for i := 999; i >= 100; i-- {
	//	x := n%i
	//	if x == 0 {
	//		fmt.Println(i, n/i, "-----")
	//	}
	//}


	//fmt.Println(strings.TrimSpace("  42"))
	//fmt.Println(leetcode.CountNumbersWithUniqueDigits(5))

	//numCourse := 4
	//prerequisites := [][]int{{1,0}, {2,0}, {3,1}, {3,2}, {4,2}}
	//prerequisites := [][]int{{3,0},{0,1}}
	//res := leetcode.FindOrder(numCourse, prerequisites)
	//fmt.Println(res)
	//graph := [][]int{{1,2},{3},{3},{}}
	//
	//fmt.Println(leetcode.AllPathsSourceTarget(graph))
	//fmt.Println(leetcode.ReachingPoints(1,13,999999989,13))

	//2,3
	//2,5
	//2,7
	//2,9
	//
	//3,3
	//3,6
	//3,9
	//12,9
	//
	//3,3
	//6,3
	//3,9

	//nums := []int{3,2,5,1}
	//nums := []int{2,3}
	//nums := []int{3,2,1,4,5,7,10,8,6,9}
	//res := basic.MergeSort(nums)
	//fmt.Println(res)


	//leetcode.SortArray(nums)
	//leetcode.QuickSort(nums, 0 , 5)
	//fmt.Println(nums)
	//
	//basic.QuickSort2(nums, 0, 5)
	//fmt.Println(nums)

    //4,2,1,3,8,7,5,6
    //head := &util.ListNode{
    //    Val: 4,
    //}
    //
    //head.Next = &util.ListNode{
    //    Val: 2,
    //}
    //
    //head.Next.Next = &util.ListNode{
    //    Val: 1,
    //}
    //
    //head.Next.Next.Next = &util.ListNode{
    //    Val: 3,
    //}
    //head.Next.Next.Next.Next = &util.ListNode{
    //    Val: 8,
    //}
    //head.Next.Next.Next.Next.Next = &util.ListNode{
    //    Val: 7,
    //}
    //
    //head.Next.Next.Next.Next.Next.Next = &util.ListNode{
    //    Val: 5,
    //}
    //head.Next.Next.Next.Next.Next.Next.Next = &util.ListNode{
    //    Val: 6,
    //}
    //
    //res := leetcode.SortList(head)
    //for res != nil {
    //    fmt.Println(res.Val)
    //    res = res.Next
    //}

    //res := leetcode.InsertionSortList(head)
    //for res != nil {
    //    fmt.Println(res.Val)
    //    res = res.Next
    //}

    //[12,1],[2],[15,11],[5,2],[1,15],[4,2],[5],[15,15]]
    //cache := leetcode.ConstructorLRU(2)
    //
    //fmt.Println(cache.Get(6))
    //fmt.Println(cache.Get(8))
    //cache.Put(12,1)
    //fmt.Println(cache.Get(2))
    //cache.Put(15,11)
    ////cache.Add(head)
    //cache.Put(5,2)
    //fmt.Println("-------")
    //fmt.Println(cache.Tail.Val)
    //cache.Put(1,15)
    //fmt.Println(cache.Tail.Val)
    //
    //cache.Put(4,2)
    //fmt.Println("-------")
    //fmt.Println(cache.Get(5))
    //
    //cache.Put(15,15)

    //cache.Remove(tail)
    //fmt.Println("-------")
    //fmt.Println(cache.Head.Val)
    //fmt.Println(cache.Tail.Val)
    //
    //fmt.Println(len(cache.Keys))

    //nums := []int{-2,3,-4}
    //fmt.Println(leetcode.MaxProduct(nums))

    //head := &util.ListNode{Val: 3}
    //head.Next = &util.ListNode{Val: 2}
    //head.Next.Next = &util.ListNode{Val: 0}
    //head.Next.Next.Next = &util.ListNode{Val: 4}
    //head.Next.Next.Next.Next = head.Next
    //res := leetcode.DetectCycle(head)
    //fmt.Println(res.Val)
    //s := "leetcode"
    //wordDict := []string{"leet", "code"}
    //res := leetcode.WordBreak(s, wordDict)
    //fmt.Println(res)

    //fmt.Println(0^3)
    //fmt.Println("----")
    //nums := []int{2,2,3,4,4}
    //fmt.Println(leetcode.SingleNumber(nums))

    //nums := []int{100,4,200,1,3,2}
    //fmt.Println(leetcode.LongestConsecutive(nums))

	/**
	      5
	    3    7
	  1  4  6  8
	 */

    //root := &util.TreeNode{
    //   Val: 5,
    //}
    //
    //root.Left = &util.TreeNode{
    //   Val: 3,
    //}
    //
    //root.Right = &util.TreeNode{
    //   Val: 7,
    //}
    //root.Left.Left = &util.TreeNode{
    //   Val: 1,
    //}
    //root.Left.Right = &util.TreeNode{
    //   Val: 4,
    //}
    //root.Right.Left = &util.TreeNode{
    //   Val: 6,
    //}
    //root.Right.Right = &util.TreeNode{
    //   Val: 8,
    //}
    //
    //res := leetcode.LevelOrderBinaryTree(root)
	//fmt.Printf("%+v", res)

    //res := basic.Tree2List(root)
    //
    //for res != nil {
    //    fmt.Println(res.Val)
    //    res = res.Right
    //}


    //nums := [][]int{}
	//nums = append(nums, []int{1,2})
	//nums = append(nums, []int{1,4})
	//nums = append(nums, []int{3,4})
	//leetcode.RemoveCoveredIntervals(nums)

	//fmt.Println(leetcode.CountPrimes(499979))
	//word1 := "abcd"
	//word2 := "bce"
	//res := leetcode.MinDistance(word1, word2)
	//fmt.Println(res)

    //nums := []int{0,1,1,2,0,1,0,2}
    //leetcode.SortColors(nums)
    //fmt.Println(nums)

    //root := &util.TreeNode{
    //    Val: 1,
    //}
    //root.Left = &util.TreeNode{
    //    Val: 2,
    //}
    //root.Right = &util.TreeNode{
    //    Val: 3,
    //}
    //
    //root.Right.Right = &util.TreeNode{
    //    Val: 4,
    //}
    //
    //basic.PrintTree2(root)

    //s1 := "hhhelworldhhello"
    //s2 := "hhell"
    //
    //fmt.Println(leetcode.StrStr(s1, s2))

    //candi := []int{10,1,2,7,6,1,5}
    //fmt.Println(leetcode.CombinationSum40(candi, 8))


    //
    //fmt.Println(nums)

    //nums := []int{7,7}
    //res := leetcode.SearchRange(nums, 7)
    //fmt.Println(res)

    //nums := []int{4,5,6,7,0,1,2}
    //res := leetcode.Search33_2(nums, 4)
    //fmt.Println(res)

    //s := "()(()"
    //
    //fmt.Println(leetcode.LongestValidParentheses(s))

    //nums := []int{5,4,7,5,3,2}
    //leetcode.NextPermutation(nums)
    //fmt.Println(nums)

    //fmt.Println(leetcode.GenerateParenthesis(3))
    //
    //s := "abcde"
    //fmt.Println(s[0:1])
    //
    //for _, i := range s {
    //    fmt.Printf("%c ", i)
    //}

	//a := make(map[int]int,3)
	//a[1] = 3
	//fmt.Println(cap(a))

	//res := leetcode.NumTrees(4)
	//fmt.Println(res)

	//rr1 := &util.SpecialTreeNode{
	//	Val: 7,
	//}
	//
	//rl1 := &util.SpecialTreeNode{
	//	Val: 6,
	//}
	//
	//lr1 := &util.SpecialTreeNode{
	//	Val: 5,
	//}
	//
	//ll1 := &util.SpecialTreeNode{
	//	Val: 4,
	//}
	//
	//r1 := &util.SpecialTreeNode{
	//	Val: 3,
	//	Left: rl1,
	//	Right: rr1,
	//}
	//l1 := &util.SpecialTreeNode{
	//	Val: 2,
	//	Left: ll1,
	//	Right: lr1,
	//}
	//root := &util.SpecialTreeNode{
	//	Val: 1,
	//	Right: r1,
	//	Left: l1,
	//}

	//preorder := []int{1,2,4,5,3,6,7}
	//postorder := []int{4,5,2,6,7,3,1}
	//preorder := []int{2,1}
	//postorder := []int{1,2}
	//res := leetcode.ConstructFromPrePost(preorder, postorder)
	//fmt.Println(res.Val)
	//fmt.Println(res.Left.Val, res.Right.Left.Val, res.Left.Right.Val)

	//
	//res := leetcode.Connect(root)
	//
	//fmt.Println(res.Left.Next.Val, res.Left.Left.Next.Val, res.Left.Right.Next.Val)

	//nums := []int{1,3,-1,-3,5,3,6,7}
	//res := leetcode.MaxSlidingWindow(nums, 3)
	//fmt.Println(res)

	//nums := []int{73,74,75,71,69,72,76,73}
	//res := leetcode.DailyTemperatures(nums)
	//fmt.Println(res)

	//nums1 := []int{4,1,2}
	//nums2 := []int{1,3,4,2}
	//res := leetcode.NextGreaterElement(nums1, nums2)
	//fmt.Println(res)

	//nums := []int{0,1,0,3,12}
	//leetcode.MoveZeroes(nums)


	//nums := []int{1}
	//res := leetcode.RemoveElement(nums, 1)
	//fmt.Println(res)

	//nums := []int{1,2,3,1,1}
	//res := leetcode.ShipWithinDays(nums, 4)
	//fmt.Println(res)

	//nums := []int{-1,0,3,3,5,9,12}
	//res := leetcode.BinarySearch2(nums, 3)
	//fmt.Println(res)
	//s := "abab"
	//res := leetcode.LengthOfLongestSubstring4(s)
	//fmt.Println(res)

	//s := "abab"
	//p := "ab"
	//res := leetcode.FindAnagrams(s, p)
	//fmt.Println(res)

	//s1 := "ab"
	//s2 := "eidboaooo"
	//res := leetcode.CheckInclusion(s1, s2)
	//fmt.Println(res)

	//s := "ADOBECODEBANC"
	//t := "BC"
	//
	//res := leetcode.MinWindow2(s, t)
	//fmt.Println(res)


	//bookings := [][]int{{3,2,7},{3,7,9},{8,3,9}}
	//res := leetcode.CarPooling(bookings, 11)
	//fmt.Println(res)

	//bookings := [][]int{{1,2,10},{2,2,15}}
	//res := leetcode.CorpFlightBookings(bookings, 2)
	//fmt.Println(res)

	//nums := []int{1,1,1}
	//res := leetcode.SubarraySum2(nums, 2)
	//fmt.Println(res)

	//obj := leetcode.Constructor(nums)
	//
	//res := obj.SumRange(2,5)

	//bits := []int{1,1,0}
	//res := leetcode.IsOneBitCharacter(bits)
	////res := leetcode.GF_IsOneBitCharacter(bits)
	//fmt.Println(res)

	l1 := &util.ListNode{
		Val: 1,
	}
	l1.Next = &util.ListNode{
		Val: 2,
	}
	l1.Next.Next = &util.ListNode{
		Val: 3,
	}
	l1.Next.Next.Next = &util.ListNode{
		Val: 4,
	}
	l1.Next.Next.Next.Next = &util.ListNode{
		Val: 5,
	}
    //res:= leetcode.IsPalindromeList(l1)
    //fmt.Println(res)


    //res := leetcode.ReverseList(l1)
	//fmt.Println("--")
	//for res != nil {
	//	fmt.Println(res.Val)
	//	res = res.Next
	//}

	res := leetcode.ReverseBetween2(l1, 1,5)
	fmt.Println("--")
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	//for l1 != nil {
	//	fmt.Println(l1.Val)
	//	l1 = l1.Next
	//}


	//res:= leetcode.MajorityElement([]int{1,2,2,3,3,3,2,2,2})
	//fmt.Println(res)

	//l1 := &util.ListNode{
	//	Val: 2,
	//}
	//l1.Next = &util.ListNode{
	//	Val: 4,
	//}
	//l1.Next.Next = &util.ListNode{
	//	Val: 6,
	//}
	//l1.Next.Next.Next = &util.ListNode{
	//	Val: 7,
	//}
	//l2 := &util.ListNode{
	//	Val: 1,
	//}
	//l2.Next = &util.ListNode{
	//	Val: 3,
	//}
	//l2.Next.Next = &util.ListNode{
	//	Val: 5,
	//}
	//
	//res:= leetcode.MergeTwoLists2(l1, l2)
	//for res != nil {
	//	fmt.Println(res.Val)
	//	res = res.Next
	//}

	//res:= leetcode.Permute([]int{1,2,3})
	//fmt.Println(res)

	//res:= leetcode.LetterCombinations("22")
	//fmt.Println(res)

	//res := leetcode.IsValid("){")
	//fmt.Println(res)

	//res := leetcode.MaxProfit([]int{7,1,5,3,6,4})
	//fmt.Println(res)

	//r2 := &util.TreeNode{
	//	Val: 4,
	//}
	//l2 := &util.TreeNode{
	//	Val: 3,
	//}
	//r1 := &util.TreeNode{
	//	Val: 3,
	//}
	//l1 := &util.TreeNode{
	//	Val: 2,
	//	//Left: l2,
	//	//Right: r2,
	//}
	//root := &util.TreeNode{
	//	Val: 1,
	//	Right: r1,
	//	Left: l1,
	//}


	//leetcode.Flatten(root)
	//fmt.Print(root.Right.Val, root.Right.Right.Val, root.Right.Right.Right.Val, )

	//res := leetcode.InvertTree(root)
	//fmt.Print(res.Left.Val, res.Left.Left.Val)
	//
	//res := leetcode.IsSymmetric(root)
	//fmt.Println(res)

	//res := leetcode.InorderTraversal(root)
	//fmt.Println(res)

	//nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	//res := leetcode.MaxSubArray(nums)
	//fmt.Println(res)
	//l1 := &util.ListNode{
	//	Val: 2,
	//}
	//l1.Next = &util.ListNode{
	//	Val: 3,
	//}
	//l1.Next.Next = &util.ListNode{
	//	Val: 4,
	//}
	//l1.Next.Next.Next = &util.ListNode{
	//	Val: 5,
	//}
	//res := leetcode.RemoveNthFromEnd(l1, 4)
	//fmt.Println(res.Val,res.Next.Val, res.Next.Next.Val)

	//fmt.Println(math.MinInt32, math.MaxInt32, math.MaxUint32)
	//res := leetcode.ReverseInt(1534236469)
	//res := leetcode.ReverseInt(1563847412)
	//fmt.Println(res)

	//res := leetcode.MaxArea2([]int{1,8,6,2,5,4,8,3,7})
	//fmt.Println(res)

	//res := leetcode.LongestPalindrome("babde")
	//fmt.Println(res)
	// 342 465
	//l1 := &util.ListNode{
	//	Val: 2,
	//}
	//l1.Next = &util.ListNode{
	//	Val: 4,
	//}
	//l1.Next.Next = &util.ListNode{
	//	Val: 3,
	//}
	//
	//l2 := &util.ListNode{
	//	Val: 5,
	//}
	//l2.Next = &util.ListNode{
	//	Val: 6,
	//}
	//l2.Next.Next = &util.ListNode{
	//	Val: 4,
	//}
	//fmt.Println((l1.Val+l2.Val)%10)
	////leetcode.AddTwoNumbers2(l1, l2)
	//newHead := leetcode.AddTwoNumbers2(l1, l2)
	//for newHead != nil {
	//	fmt.Println(newHead.Val, "--")
	//	newHead = newHead.Next
	//}
	//fmt.Println(newHead.Val)
	//fmt.Println(newHead.Next.Val)
	//fmt.Println(newHead.Next.Next.Val)
	//fmt.Println(newHead.Next.Next.Next.Val)

	//fmt.Println("---")
	//util.Max(1,2)
	//
	//nums := []int{1,3,5,7,8}
	//fmt.Println(leetcode.TwoSum(nums, 9))
}
