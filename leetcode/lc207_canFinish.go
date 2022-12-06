package leetcode

/**
构建图的邻接表
返回dfs遍历，统计onpath 和 visited

 */
func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := buildGraph(numCourses, prerequisites)
    onPath := make([]bool, numCourses)
    visited := make([]bool, numCourses)
    hasCycle := false
    var dfs func(graph [][]int, curr int)
    dfs = func(graph [][]int, curr int) {
        if onPath[curr] {
            hasCycle = true
        }
        if hasCycle || visited[curr] {
            return
        }

        visited[curr] = true
        onPath[curr] = true
        for i := 0; i < len(graph[curr]); i++ {
            dfs(graph, graph[curr][i])
        }
        onPath[curr] = false
    }
    for i := 0; i < numCourses; i++ {
        dfs(graph, i)
    }
    return !hasCycle
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
    graph := make([][]int, numCourses)
    for i := 0; i < numCourses; i++ {
        graph[i] = []int{}
    }

    for i := 0; i < len(prerequisites); i++ {
        from, to := prerequisites[i][1], prerequisites[i][0]
        graph[from] = append(graph[from], to)
    }

    return graph
}



