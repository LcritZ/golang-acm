package leetcode

/**
深度遍历，由边上的引导扩张到连续的可到达的地方



 */

func numEnclaves(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		dfs1020(grid, i, 0)
		dfs1020(grid, i, m-1)
	}
	for j := 0; j < m; j++ {
		dfs1020(grid, 0, j)
		dfs1020(grid, n-1, j)
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}

	return res
}

func dfs1020(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] != 1 {
		return
	}
	grid[i][j] = 2
	dfs1020(grid, i+1, j)
	dfs1020(grid, i-1, j)
	dfs1020(grid, i, j+1)
	dfs1020(grid, i, j-1)
}

//
//func numEnclaves(grid [][]int) int {
//	if len(grid) == 0 {
//		return 0
//	}
//	//把周围一圈改成和岛屿一样的，连在一起，这样就能搞出封闭的岛屿 封闭岛屿=岛屿-边上的岛  把边上的岛排除点，剩下的就是封闭了
//	n, m := len(grid), len(grid[0])
//	for i := 0; i<len(grid); i++ {
//		dfs1020(grid, i, 0)
//		dfs1020(grid,i,m-1)
//	}
//	for j := 0; j<m; j++ {
//		dfs1020(grid, 0, j)
//		dfs1020(grid,n-1,j)
//	}
//	res := 0
//	for i := 0; i < n; i++ {
//		for j := 0; j < m; j++ {
//			if grid[i][j] == 1 {
//				res++
//			}
//		}
//	}
//	return res
//
//}
//
//func dfs1020(grid [][]int, i, j int) {
//	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
//		return
//	}
//	if grid[i][j] != 1 {
//		return
//	}
//	grid[i][j] = 2
//	dfs1020(grid, i+1, j)
//	dfs1020(grid, i-1, j)
//	dfs1020(grid, i, j+1)
//	dfs1020(grid, i, j-1)
//}
//
