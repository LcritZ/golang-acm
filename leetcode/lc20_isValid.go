package leetcode

import "fmt"

func IsValid(str string) bool {
	if len(str) == 0 || len(str) == 1 || len(str)/2 == 1 {
		return false
	}
	charArr := []int32{}
	for _, char := range str {
		if char == '(' || char == '{' || char == '[' {
			charArr = append(charArr, char)
			continue
		}
		switch char {
		case ')':
			if charArr[len(charArr)-1] == '(' {
				charArr = charArr[:len(charArr)-1]
			} else {
				fmt.Println(")--")
				return false
			}
		case '}':
			if charArr[len(charArr)-1] == '{' {
				charArr = charArr[:len(charArr)-1]
			} else {
				fmt.Println("}--")
				return false
			}
		case ']':
			if charArr[len(charArr)-1] == '[' {
				charArr = charArr[:len(charArr)-1]
			} else {
				fmt.Println("]--")
				return false
			}

		}
	}

	return len(charArr) == 0
}
