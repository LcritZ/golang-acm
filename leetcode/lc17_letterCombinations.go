package leetcode

var letterMap = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

var result []string

func LetterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	result = []string{}
	traceLetter(digits, 0, "")
	return result
}

func traceLetter(digits string, index int, item string) {
	if len(item) == len(digits) {
		result = append(result, item)
	} else {
		num := digits[index]
		letters := letterMap[num]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			traceLetter(digits, index + 1, item + string(letters[i]))
		}
	}
}
