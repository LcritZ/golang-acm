package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {

    res := [][]string{}

    strMap := make(map[string][]string)
    for _, str := range strs {
        s := []byte(str)
        sort.Slice(s, func(i, j int) bool {
            return s[i] < s[j]
        })
        newS := string(s)
        strMap[newS] = append(strMap[newS], str)
    }

    for _, v := range strMap {
        res = append(res, v)
    }

    return res
}

func groupAnagrams2(strs []string) [][]string {

    res := [][]string{}
    charMap := make(map[[26]int][]string)

    for i := 0; i < len(strs); i++ {
        charArr := [26]int{}
        for _, c := range strs[i] {
            charArr[c-'a']++
        }
        charMap[charArr] = append(charMap[charArr], strs[i])
    }

    for _, v := range charMap {
        res = append(res, v)
    }

    return res
}
