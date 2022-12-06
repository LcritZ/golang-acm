package leetcode

/**
abbcc
bbcca
bccab

判断子串是否符合要求
那么在s中找到第一个goal首字母，s从这个index的子串要等于goal开头相同长度的子串，剩下同样要相等。不满足则找到下一个，直到遍历完成

 */
func rotateString(s string, goal string) bool {
    if len(s) != len(goal) || len(s) == 0{
        return false
    }

    first := goal[0]
    for i := 0; i < len(s); i++ {
        if s[i] == first {
            subStr := s[i:]
            length := len(subStr)
            if s[0:i]== goal[length:] && s[i:]==goal[0:length] {
                return true
            }
        }
    }

    return false
}
