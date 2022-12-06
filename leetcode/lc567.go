package leetcode

func checkInclusion(s1 string, s2 string) bool {
	charMap := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		charMap[s1[i]]++
	}
	currMap := make(map[byte]int)
	for key, _ := range charMap {
		currMap[key] = 0
	}

	left, right := 0, 0
	valid := false

	for right <= len(s2) {
		for valid {
			if right-left == len(s1) {
				return true
			} else {
				if _, ok := currMap[s2[left]]; ok {
					currMap[s2[left]]--
					if currMap[s2[left]] < charMap[s2[left]] {
						valid = false
					}
				}
				left++
			}
		}
		if right < len(s2) {
			if _, ok := charMap[s2[right]]; ok {
				currMap[s2[right]]++
			}
			flag := true
			for key, _ := range charMap {
				if currMap[key] < charMap[key] {
					flag = false
					break
				}
			}
			valid = flag
		}
		right++
	}

	return false
}
