package leetcode

/**
要理解题目的意思，聚合在一块的，周围是海那说明是岛屿，不是说一个1，而是连起来的一群1

思路：深度遍历
 */
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs200(grid, i, j)
			}
		}
	}
	return res
}

func dfs200(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	dfs200(grid, i+1, j)
	dfs200(grid, i-1, j)
	dfs200(grid, i, j+1)
	dfs200(grid, i, j-1)
} 
