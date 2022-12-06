package leetcode

func FindAnagrams(s string, p string) []int {
    res := []int{}

    if len(p) > len(s) {
        return res
    }
    charMap := make(map[byte]int)
    for i := 0; i < len(p); i++ {
        charMap[p[i]]++
    }
    currMap := make(map[byte]int)
    for key, _ := range charMap {
        currMap[key] = 0
    }

    left, right := 0, 0
    valid := false

    for right <= len(s) {
        for valid {
            //判断是否相等，相等则说明可以加入到结果，然后跳出来，让right继续滑动
            if right-left == len(p) {
                res = append(res, left)
                break
            } else {
                if _, ok := currMap[s[left]]; ok {
                    currMap[s[left]]--
                    if currMap[s[left]] < charMap[s[left]] {
                        valid = false
                    }
                }
            }
            left++
        }
        if right < len(s) {
            if _, ok := charMap[s[right]]; ok {
                currMap[s[right]]++
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

    return res
}
