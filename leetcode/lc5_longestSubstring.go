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

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		s1, e1 := palindromeHelper(s, i,  i)
		s2, e2 := palindromeHelper(s, i, i+1)
		if e1-s1 > end-start {
			start, end = s1, e1
		}
		if e2-s2 > end-start {
			start, end = s2, e2
		}
	}
	fmt.Println(start, end)
	return s[start:end+1]
}

func palindromeHelper(s string, start, end int) (int, int) {
	for ;start>=0 && end<len(s) && s[start] == s[end]; {
		start, end = start-1, end+1
	}
	return start+1, end-1
}


func GF_longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i + 1)
		if right1 - left1 > end - start {
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1 , right+1 { }
	return left + 1, right - 1
}
