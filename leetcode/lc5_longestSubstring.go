package leetcode

import "fmt"

/**
最长回文字符串
 */
func LongestSubstring(s string) string {

	if len(s) == 0 || len(s) == 1 {
		return s
	}
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		l1, r1 := substringHelper(s, i, i)
		l2, r2 := substringHelper(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
		fmt.Println(start, end)
	}
	return s[start:end+1]
}

func substringHelper(s string, left, right int) (int, int) {
	//轮询
	for ; left >= 0 && right < len(s) && s[left] == s[right]; {
		left, right = left-1, right+1
	}
	//回退
	return left+1, right-1
}
