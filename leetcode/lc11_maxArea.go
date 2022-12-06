package leetcode

import (
	"fmt"
	"golang-acm/util"
)

/**
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。
在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。
输入：[1,8,6,2,5,4,8,3,7]
输出：49


移动指针是看怎样子能够拿到更大的值，哪个小说明移动另一边是无用的，必须移动小的一方

*/
//时间复杂度高
func MaxArea1(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var max int
	for i := 0; i < len(height)-1; i++ {
		if i > 0 && height[i] <= height[i-1] {
			continue
		}
		for j := i+1; j < len(height); j++ {
			fmt.Println(i, j, "--")
			h := util.Min(height[i], height[j])
			if h*(j-i) > max {
				max = h*(j-i)
			}
		}
	}

	return max
}

//type StartEndIndex struct {
//	start int
//	end   int
//}

func MaxArea2(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}
	i := 0
	j := len(height)-1

	high := util.Min(height[i], height[j])
	max := high*(j-i)

	for i < j {
		if height[i] <= height[j] {
			high = height[i]
			max = util.Max(max, high*(j-i))
			i++
		} else {
			high = height[j]
			max = util.Max(max, high*(j-i))
			j--
		}
	}

	return max

}

func MaxArea3(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	left, right := 0, len(height)-1
	max := 0
	for left < right {
		h := util.Min(height[left], height[right])
		max = util.Max(max, h*(right-left))
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return max

}