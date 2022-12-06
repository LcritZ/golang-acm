package leetcode

import "strconv"

/**
输入：ops = ["5","-2","4","C","D","9","+","+"]
输出：27
解释：
"5" - 记录加 5 ，记录现在是 [5]
"-2" - 记录加 -2 ，记录现在是 [5, -2]
"4" - 记录加 4 ，记录现在是 [5, -2, 4]
"C" - 使前一次得分的记录无效并将其移除，记录现在是 [5, -2]
"D" - 记录加 2 * -2 = -4 ，记录现在是 [5, -2, -4]
"9" - 记录加 9 ，记录现在是 [5, -2, -4, 9]
"+" - 记录加 -4 + 9 = 5 ，记录现在是 [5, -2, -4, 9, 5]
"+" - 记录加 9 + 5 = 14 ，记录现在是 [5, -2, -4, 9, 5, 14]
所有得分的总和 5 + -2 + -4 + 9 + 5 + 14 = 27

 */
func calPoints(ops []string) int {
    if len(ops) == 0 {
        return 0
    }

    res := 0
    points := []int{}
    for i := 0; i < len(ops); i++ {
        if ops[i] == "C" {
            points = points[:len(points)-1]
        } else if ops[i] == "D" {
            points = append(points, points[len(points)-1]*2)
        } else if ops[i] == "+" {
            points = append(points, points[len(points)-1]+points[len(points)-2])
        } else {
            p, _ := strconv.Atoi(ops[i])
            points = append(points, p)
        }
    }

    for i := 0; i < len(points); i++ {
        res += points[i]
    }

    return res
}

