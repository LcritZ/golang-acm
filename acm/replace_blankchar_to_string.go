package acm

import (
	"unicode/utf8"
)

func ReplaceSpace( s string ) string {
	if utf8.RuneCountInString(s) == 0 {
		return s
	}
	var res string
	for _, str := range s {
		if string(str) == " " {
			res += "%20"
		} else {
			res += string(str)
		}
	}

	//第二种解法
	//strings.ReplaceAll(s, " ", "%20")

	return res
}
