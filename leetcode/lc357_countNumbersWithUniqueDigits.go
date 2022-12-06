package leetcode

/**
n==3
0-999

f0 = 1
f1 = 1*10
f2 = 10 + 前部分可选(f1-f0)*(11-n)后部分可选 = 91

f3 = f2 + (f2-f1)*(11-3) = 91+ 81*8 = 739

100-9-10

f3 = 90*
f(n-1)*(11-n)+

 */
func CountNumbersWithUniqueDigits(n int) int {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return 10
    }

    x, y := 1, 10
    for i := 2; i <= n; i++ {
        temp := y+(y-x)*(11-i)
        x, y = y, temp
    }

    return y
}

