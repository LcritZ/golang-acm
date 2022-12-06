package leetcode

import (
    "fmt"
)

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

输入：s = "ADOBECODEBANC", t = "BC"
输出："BANC"

s="abcd"
t="ac"

思路：滑动窗口，对T中字符用map保存，然后顺序遍历S, 出现一个保存一个，直到全部覆盖，记录一下当前长度，再滑动

需要额外的map存储，并且需要存储当前是否有效，并在移动left, right的时候判断，修改状态

注意取值right与right++的关系，如果是在++前使用，就需要+1，
 */
func MinWindow(s string, t string) string {
    charMap := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        charMap[t[i]]++
    }
    currMap := make(map[byte]int)
    for key, _ := range charMap {
        currMap[key] = 0
    }

    left, right := 0, 0
    res := ""
    valid := false
    //输入：s = "ADOBECODEBANC", t = "BC"
    for right <= len(s) {
        fmt.Println("curr1:", s[left:right], "res: ", res, valid)
        for valid {
            //对所有满足条件的都要判断长度是否最短，更短则更新
            if len(res) == 0 {
                res = s[left:right]
            }
            if len(res) > 0 && right-left < len(res) {
                res = s[left:right]
            }
            if _, ok := charMap[s[left]]; ok {
                currMap[s[left]]--
                if currMap[s[left]] < charMap[s[left]] {
                    valid = false
                }
            }
            left++
        }
        fmt.Println("curr2:", s[left:right], "res: ", res, valid)
        if right < len(s) {
            if _, ok := charMap[s[right]]; ok {
                currMap[s[right]]++
                if !valid {
                    flag := true
                    for key, _ := range charMap {
                        if charMap[key] > currMap[key] {
                            flag = false
                        }
                    }
                    valid = flag
                }
                fmt.Println("valid:", valid, currMap)
            }
        }
        right++

    }

    return res
}

func MinWindow2(s string, t string) string {
    charMap := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        charMap[t[i]]++
    }

    currMap := make(map[byte]int)
    for k, _ := range charMap {
        currMap[k] = 0
    }

    left, right := 0, 0
    window := ""
    valid := false
    for ;right <= len(s); right++ {
        if right < len(s) {
            flag := true
            currMap[s[right]]++
            for k, v := range charMap {
                if currMap[k] < v {
                    flag = false
                }
            }
            valid = flag
        }
        for valid {
            if len(window) == 0 {
                window = s[left:right+1]
            }
            if len(window) > 0 && len(window) > right+1-left {
                window = s[left:right+1]
            }
            currMap[s[left]]--
            if currMap[s[left]] < charMap[s[left]] {
                valid = false
            }
            left++
        }

    }

    return window
}
