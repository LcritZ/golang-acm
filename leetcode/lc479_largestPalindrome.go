package leetcode

import "math"

//99*99 = 9801
//
//9 * 9 = 81
//
//999*999 = 990099
//


//func LargestPalindrome(n int) int {
//
//}

func GF_LargestPalindrome(n int) int {
    if n == 1 {
        return 9
    }
    upper := int(math.Pow10(n)) - 1
    for left := upper; ; left-- { // 枚举回文数的左半部分
        p := left
        for x := left; x > 0; x /= 10 {
            p = p*10 + x%10 // 翻转左半部分到其自身末尾，构造回文数 p
        }
        for x := upper; x*x >= p; x-- {
            if p%x == 0 { // x 是 p 的因子
                return p % 1337
            }
        }
    }
}
