package acm

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。


示例 1：

输入：s = "()"
输出：true
示例2：

输入：s = "(]"
输出：false
示例4：

输入：s = "([)]"
输出：false
示例5：

 */

func IsSeq(str string) bool {
	if len(str) == 0 || len(str) == 1  {
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
				return false;
			}
		case '}':
			if charArr[len(charArr)-1] == '{' {
				charArr = charArr[:len(charArr)-1]
			} else {
				return false;
			}
		case ']':
			if charArr[len(charArr)-1] == '(' {
				charArr = charArr[:len(charArr)-1]
			} else {
				return false;
			}

		}
		//if char == ')' {
		//	if charArr[len(charArr)-1] == '(' {
		//
		//	} else {
		//
		//	}
		//}
	}


	return false
}

