package leetcode

func isBipartite(graph [][]int) bool {
    length := len(graph)
    visited := make([]bool, length)
    color := make([]bool, length)

    res := true

    var dfs func(graph [][]int, index int)
    dfs = func(graph [][]int, index int)  {
        if !res {
            return
        }
        visited[index] = true
        for j := 0; j < len(graph[index]); j++ {
            if !visited[graph[index][j]] {
                visited[graph[index][j]] = true
                color[graph[index][j]] = !color[index]
                dfs(graph, graph[index][j])
            }  else {
                if color[graph[index][j]] == color[index] {
                    res = false
                    return
                }
            }
        }
    }

    // color[0] = true
    for i := 0; i < len(graph) && res; i++ {
        if !visited[i] {
            color[i] = true
            dfs(graph, i)
        }
    }

    return res
}
