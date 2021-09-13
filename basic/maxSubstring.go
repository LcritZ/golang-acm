package basic

func LongestPalindrome(s string) string {
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
	for ; left >= 0 && right < len(s) && s[left] == s[right]; {
		left, right = left-1 , right+1
	}
	return left + 1, right - 1
}


func t2(s string) string {

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)
		if r1 - l1 > end - start {
			start, end = l1, r1
		}
		if r2 - l2 > end - start {
			start, end = l2, r2
		}
	}
	return s[start:end+1]
}

func e2(s string, l, r int) (int, int) {
	for ; l >= 0 && r < len(s) && s[l] == s[r]; {
		l ,r = l-1, r+1
	}
	return l+1, r-1
}