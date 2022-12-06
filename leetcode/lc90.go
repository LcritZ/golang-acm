package leetcode

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	k := len(nums)
	sort.Ints(nums)

	var trace func(subArr []int, index, length int)
	trace = func(subArr []int, index, length int) {
		if len(subArr) == length {
			temp := make([]int, len(subArr))
			copy(temp, subArr)
			res = append(res, temp)
			return
		}
		for i := index+1; i < len(nums); i++ {
			if i > index+1 && i < len(nums) && nums[i] == nums[i-1] {
				continue
			}
			subArr = append(subArr, nums[i])
			trace(subArr, i, length)
			subArr = subArr[0:len(subArr)-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		subArr := []int{nums[i]}
		for j := 1; j <= k-i; j++ {
			trace(subArr, i, j)
		}
	}

	return res
}

