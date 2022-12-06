package leetcode

func AllPathsSourceTarget(graph [][]int) [][]int {
    if len(graph) == 0 {
        return [][]int{}
    }

    res := [][]int{}
    var dfs func(graph [][]int, curr []int, index int)
    dfs = func(graph [][]int, curr []int, index int) {
        if index == len(graph)-1 {
            return
        }
        temp := graph[index]
        for i := 0; i < len(temp); i++ {
            if temp[i] == len(graph)-1 {
                temp1 := append(curr, temp[i])
                res = append(res, append([]int{}, temp1...))
            } else {
                dfs(graph, append(curr, temp[i]), temp[i])
            }
        }
    }
    curr := []int{0}
    dfs(graph, curr, 0)

    return res

}
