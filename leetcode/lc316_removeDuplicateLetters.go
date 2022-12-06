package leetcode

/**

给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

输入：s = "bcabc"
输出："abc"

输入：s = "cbacdcbc"
输出："acdb"

思路：用26个字母数组保存字母出现次数
然后遍历数组，找到按顺序的第一个字母=1的情况，然后拼接的字母即可

还要考虑相对的位置

 */
//func RemoveDuplicateLetters(s string) string {
//    if len(s) == 0 {
//        return s
//    }
//
//    wordMap := make(map[byte]int)
//    for i := 0; i < len(s); i++ {
//        wordMap[s[i]]++
//    }
//
//    wordArr := make([]byte, len(wordMap))
//
//    for i := 0; i < len(s); i ++ {
//
//    }
//
//
//}

func GF_removeDuplicateLetters(s string) string {
    left := [26]int{}
    for _, ch := range s {
        left[ch-'a']++
    }
    stack := []byte{}
    inStack := [26]bool{}
    for i := range s {
        ch := s[i]
        if !inStack[ch-'a'] {
            for len(stack) > 0 && ch < stack[len(stack)-1] {
                last := stack[len(stack)-1] - 'a'
                if left[last] == 0 {
                    break
                }
                stack = stack[:len(stack)-1]
                inStack[last] = false
            }
            stack = append(stack, ch)
            inStack[ch-'a'] = true
        }
        left[ch-'a']--
    }
    return string(stack)
}
