package leetcode

import "fmt"

/**
KMP字符串匹配

haystack = "lbhelalo",

needle = "lalo"
pi[1] = 0
i = 2, j = 1
pi[2] = 1
pi[3] = 1

 */
func StrStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			fmt.Printf("---, %d, %c, %c\n", j, needle[i], needle[j])
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
		fmt.Println(i, j, pi)
	}
	for i, j := 0, 0; i < n; i++ {
		fmt.Println("i--j", i, j)
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func StrStr2(haystack, needle string) int {

	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	next := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j >0 && needle[i] != needle[j] {
			j = next[j-1]
		}

		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i-m+1
		}
	}

	return -1
}