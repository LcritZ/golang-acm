package leetcode

func ClimbStairs(n int) int {

    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    res := 0
    i := 3
    x,y := 1,2
    for i <=n {
        res = x+y
        x, y = y, res
        i++
    }

    return res
}

