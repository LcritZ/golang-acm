package leetcode

/**
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").

也就是说是否包含全部字符

思路：滑动窗口，移动，存储，更新状态
注意边界条件

 */
func CheckInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    charMap := make(map[byte]int)
    for i := 0; i < len(s1); i++ {
        charMap[s1[i]]++
    }
    currMap := make(map[byte]int)
    for key, _ := range charMap {
        currMap[key] = 0
    }

    left, right := 0, 0
    valid := false

    for right <= len(s2) {
        for valid {
            //判断长度相等，说明存在这个排列，可以理解返回了
            if right-left == len(s1) {
                return true
            } else {
                if _, ok := currMap[s2[left]]; ok {
                    currMap[s2[left]]--
                    if currMap[s2[left]] < charMap[s2[left]] {
                        valid = false
                    }
                }
            }
            left++
        }
        if right < len(s2) {
            if _, ok := charMap[s2[right]]; ok {
                currMap[s2[right]]++
            }
            flag := true
            for key, _ := range charMap {
                if currMap[key] < charMap[key] {
                    flag = false
                    break
                }
            }
            valid = flag
        }
        right++
    }

    return false
}
