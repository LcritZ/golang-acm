package leetcode


/**
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案
输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

回溯框架

DFS遍历
判断满足条件就加进去

然后选择可以选择的，选择完回退

添加数组一定要copy, 不然的话，复用底层数组，就会被覆盖掉
copy需要给足够的空间
 */
func Combine(n int, k int) (ans [][]int) {
	var traverse func(subArr []int)
	traverse = func(subArr []int) {
		if len(subArr) == k {
			//核心需要copy
			temp := make([]int, k)
			copy(temp, subArr)
			ans = append(ans, temp)
			return
		}
		x := subArr[len(subArr)-1]
		for i := x+1; i <= n; i++ {
			subArr = append(subArr, i)
			traverse(subArr)
			subArr = subArr[0:len(subArr)-1]
		}
	}
	for i := 1; i <= n-k+1; i++ {
		subArr := []int{i}
		traverse(subArr)
	}
	return
}

func GF_combine(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp) + (n - cur + 1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}
