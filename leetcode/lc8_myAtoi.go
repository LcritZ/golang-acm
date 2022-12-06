package leetcode

import (
    "math"
    "strings"
)

func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if len(s) == 0 {
        return 0
    }
    abs, sign, i := 0, 1, 0
    if s[i] == '-' {
        sign = -1
        i++
    } else if s[i] == '+' {
        sign = 1
        i++
    }
    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        abs = 10*abs + int(s[i]-'0')
        if sign*abs < math.MinInt32 {
            return math.MinInt32
        } else if sign*abs > math.MaxInt32 {
            return math.MaxInt32
        }
        i++
    }

    return sign*abs
}

func GF_myAtoi(s string) int {
    abs, sign, i, n := 0, 1, 0, len(s)
    //丢弃无用的前导空格
    for i < n && s[i] == ' ' {
        i++
    }
    //标记正负号
    if i < n {
        if s[i] == '-' {
            sign = -1
            i++
        } else if s[i] == '+' {
            sign = 1
            i++
        }
    }
    for i < n && s[i] >= '0' && s[i] <= '9' {
        abs = 10*abs + int(s[i]-'0')  //字节 byte '0' == 48
        if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
            return math.MinInt32
        } else if sign*abs > math.MaxInt32 {
            return math.MaxInt32
        }
        i++
    }
    return sign * abs
}

