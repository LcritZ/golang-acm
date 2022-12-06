package leetcode

func closedIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	//把周围一圈改成和岛屿一样的，连在一起，这样就能搞出封闭的岛屿 封闭岛屿=岛屿-边上的岛  把边上的岛排除点，剩下的就是封闭了
	n, m := len(grid), len(grid[0])
	for i := 0; i<len(grid); i++ {
		dfs1254(grid, i, 0)
		dfs1254(grid,i,m-1)
	}
	for j := 0; j<m; j++ {
		dfs1254(grid, 0, j)
		dfs1254(grid,n-1,j)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				res++
				dfs1254(grid, i, j)
			}
		}
	}
	return res

}

func dfs1254(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] != 0 {
		return
	}
	grid[i][j] = 2
	dfs1254(grid, i+1, j)
	dfs1254(grid, i-1, j)
	dfs1254(grid, i, j+1)
	dfs1254(grid, i, j-1)
}

