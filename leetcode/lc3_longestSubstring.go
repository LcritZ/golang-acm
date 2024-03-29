package leetcode

import (
	"fmt"
	"golang-acm/util"
)

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
abcbefb

dp[i][j] = dp[i][j-1] +1  dp[j] != dp[j-1]
dp[i][j] = dp[i][j-1]

滑动窗口，重要的处理点是边界，也就是left移动不能直接跳过，需要一步一步，用map存位置可能被right更新掉了

移动left, 需要更新存的值
 */
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	ans := 0
	j := 0
	charMap := map[byte]int{}
	for i := 0; i < len(s); {
		for ; j < len(s) && charMap[s[j]] == 0; j++ {
			charMap[s[j]]++
		}
		ans = util.Max(ans, j-i)
		if j == len(s) {
			break
		}
		charMap[s[i]]--
		i++
	}

	return ans
}

func GF_LengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	ans := 0
	j := 0
	charMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		for ; j < len(s) && charMap[s[j]] == 0; j++ {
			charMap[s[j]]++
		}
		ans = util.Max(ans, j-i)
		if j == len(s) {
			break
		}
		charMap[s[i]]--
	}

	return ans
}

func LengthOfLongestSubstring2(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	charCountMap := make([]int, 256) // 字符词频表
	r := 0
	ans := 0
	for l := 0; l < n; l++ {
		fmt.Println(charCountMap)
		for ; r < n && charCountMap[s[r]] == 0; r++ {
			charCountMap[s[r]]++
		}
		fmt.Println(r)
		ans = util.Max(ans, r - l)
		if r == n {
			break
		}
		charCountMap[s[l]]--
	}
	return ans
}


func LengthOfLongestSubstring3(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	ans := 1
	lastIndexMap := make([]int, 256) // 字符上一次出现的位置表
	for i := 0; i < 256; i++ {
		lastIndexMap[i] = -1 // 初始：-1，标记都没出现过
	}
	// dp含义：dp[i]表示 以i位置做结尾的不重复子串的开始位置，即：[dp[i] ... i]是以i结尾时的最长不重复子串
	dp := make([]int, n)
	lastIndexMap[s[0]] = 0
	dp[0] = 0 // [0...0]不重复
	for i := 1; i < n; i++ {
		// i向前最远能推到哪里？1）i-1向前推到的位置；2）[i]上一次出现的位置的下一个位置；二者取决于瓶颈：较大者（最靠近i的）
		dp[i] = util.Max(dp[i-1], lastIndexMap[s[i]]+1)
		lastIndexMap[s[i]] = i
		ans = util.Max(ans, i - dp[i] + 1)
	}
	return ans
}

//滑动窗口的解法
func LengthOfLongestSubstring4(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	left, right := 0, 0
	//valid := true
	length := 0
	charMap := make(map[byte]int)
	for right < len(s) {
		fmt.Println(s[left:right])
		charMap[s[right]]++
		for charMap[s[right]] > 1 {
			left++
			charMap[s[left]]--
		}
		right++
		length = util.Max(length, right-left)
	}
	return length
}

func LengthOfLongestSubstring5(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	left, right := 0, 0
	maxLen := 0
	charMap := make(map[byte]int)
	for right < len(s) {
		charMap[s[right]]++
		for charMap[s[right]] > 1 {
			charMap[s[left]]--
			left++
		}
		right++
		maxLen = util.Max(maxLen, right-left)
	}

	return maxLen
}
