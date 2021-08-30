package leetcode

import (
	"fmt"
	"golang-acm/common"
)

func TwoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		fmt.Printf("%+v\n", hashTable)
		p2, _ := hashTable[11];
		fmt.Println(p2, "--")
		if p, ok := hashTable[target-x]; ok {
			fmt.Println(p,i)
			return []int{p, i}
		}
		hashTable[x] = i
	}
	common.Max(1,2)
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
