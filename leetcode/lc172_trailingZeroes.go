package leetcode

/**
0的个数转化为5的个数

 */
func GF_trailingZeroes(n int) int {
    ans := 0
    for i := 5; i <= n; i += 5 {
        for x := i; x%5 == 0; x /= 5 {
            ans++
        }
    }
    return ans
}

func GF_trailingZeroes2(n int) (ans int) {
    for n > 0 {
        n /= 5
        ans += n
    }
    return
}
