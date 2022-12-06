package leetcode

import "fmt"

func FindOrder(numCourses int, prerequisites [][]int) []int {
    graph := buildGraph210(numCourses, prerequisites)
    fmt.Println(graph)
    onPath := make([]bool, numCourses)
    visited := make([]bool, numCourses)
    path := []int{}
    hasCycle := false
    var dfs func(graph [][]int, curr int)
    dfs = func(graph [][]int, curr int) {
        if onPath[curr] {
            hasCycle = true
            return
        }
        if visited[curr] {
            return
        }
        //fmt.Println("path:", path)
        if len(path) == numCourses {
            return
        }
        visited[curr] = true
        onPath[curr] = true
        for i := 0; i < len(graph[curr]); i++ {
            dfs(graph, graph[curr][i])
        }
        path = append(path, curr)
        fmt.Println("path", path, visited)
        onPath[curr] = false
        //path = path[0:len(path)-1]
    }
    for i := 0; i < numCourses && len(path) < numCourses; i++ {
        if len(graph[i]) > 0 {
            dfs(graph, i)
        } else {
            if !visited[i] {
                path = append(path, i)
                visited[i] = true
            }
        }
    }

    if hasCycle {
        return []int{}
    }

    return path
}

func buildGraph210(numCourses int, prerequisites [][]int) [][]int {
    graph := make([][]int, numCourses)
    for i := 0; i < numCourses; i++ {
        graph[i] = []int{}
    }

    for i := 0; i < len(prerequisites); i++ {
        curr, prev := prerequisites[i][0], prerequisites[i][1]
        graph[curr] = append(graph[curr], prev)
    }

    return graph
}

