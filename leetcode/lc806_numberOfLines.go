package leetcode

func numberOfLines(widths []int, s string) []int {
    if len(s) == 0 {
        return []int{0,0}
    }

    lines := 1
    curr := 0
    for i := 0; i < len(s); i++ {
        c := s[i]-'a'
        if curr + widths[c] > 100 {
            lines++
            curr = widths[c]
        } else {
            curr += widths[c]
        }
    }

    return []int{lines, curr}


}
