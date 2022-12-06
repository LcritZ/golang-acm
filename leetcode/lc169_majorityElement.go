package leetcode

/**
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

 */

func MajorityElement(nums []int) int {
	numMap := make(map[int]int, 1)
	var key int
	for i := 0; i < len(nums); i++ {
		if len(numMap) == 0 {
			key = nums[i]
			numMap[key]++
			continue
		}
		if key == nums[i] {
			numMap[key]++
		} else {
			numMap[key]--
			if numMap[key] == 0 {
				delete(numMap, key)
			}
		}
	}

	return key
}


