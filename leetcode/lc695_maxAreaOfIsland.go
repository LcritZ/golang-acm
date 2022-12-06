package leetcode

import "golang-acm/util"

/**
依旧深度遍历，不过要统计每次找到岛屿后大小

 */


func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				res = util.Max(res, dfs695(grid, i, j))
			}
		}
	}
	return res
}

func dfs695(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1+dfs695(grid, i+1, j) + dfs695(grid, i-1, j) + dfs695(grid, i, j+1) + dfs695(grid, i, j-1)
}
