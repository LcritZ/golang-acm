package leetcode

func TwoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}


func TwoSum2(nums []int, target int) []int {
	//if len(nums) == 0 {
	//	return
	//}
	l := 0
	r := len(nums)-1
	for l < r {
		if nums[l] + nums[r] > target {
			l++
		} else if nums[l] + nums[r] == target {
			return []int{l, r}
		} else {
			r--
		}
	}
	return nil
}
