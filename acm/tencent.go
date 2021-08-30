package acm

/**
有效括号
({[]})

({)}

 */

func IsValid(str string) bool {
	if len(str) == 0 {
		return false
	}

	arr := []int32{}
	for _, char := range str {
		if char == '(' || char == '{' || char == '[' {
			arr = append(arr, char)
			continue
		}
		switch char {
		case ')':
			if arr[len(arr)-1] == '(' {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		case '}':
			if arr[len(arr)-1] == '{' {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		case ']':
			if arr[len(arr)-1] == '[' {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}

	if len(arr) > 0 {
		return false
	}

	return true
}