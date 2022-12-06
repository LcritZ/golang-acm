package leetcode

/**
穷举
长度从1开始到最长

 */
func Subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	k := len(nums)

	var dfs func(subArr []int, index, length int)
	dfs = func(subArr []int, index, length int) {
		if len(subArr) == length {
			temp := make([]int, len(subArr))
			copy(temp, subArr)
			res = append(res, temp)
			return
		}
		for i := index+1; i < len(nums); i++ {
			subArr = append(subArr, nums[i])
			dfs(subArr, i, length)
			subArr = subArr[0:len(subArr)-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		subArr := []int{nums[i]}
		for j := 1; j <= k-i; j++ {
			dfs(subArr, i, j)
		}
	}

	return res
}

func Subsets2(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	k := len(nums)

	var dfs func(curr []int, index, length int)
	dfs = func(curr []int, index, length int) {
		if length == k {
			res = append(res, append([]int{}, curr...))
			return
		}

		for i := index+1; i < len(nums); i++ {
			curr = append(curr, nums[i])
			dfs(curr, i, length)
			curr = curr[:len(curr)-1]
		}

	}

	for i := 0; i < k; i ++ {
		for j := 1; j <= k-i; j++ {
			curr := []int{nums[i]}
			dfs(curr, i, j)
		}
	}
	return res
}

func GF_subsets(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

func GF2_subsets(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

func subsets2(nums []int) (ans [][]int) {
	temp := []int{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			ans = append(ans, append([]int{}, temp...))
		}

		temp = append(temp, nums[index])
		dfs(index+1)
		temp = temp[:len(temp)-1]
		dfs(index+1)
	}

	return
}
