package leetcode

/**
有两种特殊字符：

第一种字符可以用一个比特0来表示
第二种字符可以用两个比特(10或11)来表示、
给定一个以 0 结尾的二进制数组bits，如果最后一个字符必须是一位字符，则返回 true

输入: bits = [1, 1, 1, 0]
输出: false
解释: 唯一的编码方式是两比特字符和两比特字符。
所以最后一个字符不是一比特字符。

思路：顺序遍历，

 */
func IsOneBitCharacter(bits []int) bool {
    for i := 0; i < len(bits); i++ {
        if i == len(bits)-1 {
            return true
        }
        if bits[i] == 0 {
            continue
        } else {
            if i == len(bits)-2 {
                return false
            }
            i++
        }
    }

    return true
}

func GF_IsOneBitCharacter(bits []int) bool {
    i, n := 0, len(bits)

    for i < n-1 {
        i += bits[i]+1
    }

    return i == n-1
}