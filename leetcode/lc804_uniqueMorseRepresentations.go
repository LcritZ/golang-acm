package leetcode

var morse = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}


/**
输入: words = ["gin", "zen", "gig", "msg"]
输出: 2
解释:
各单词翻译如下:
"gin" -> "--...-."
"zen" -> "--...-."
"gig" -> "--...--."
"msg" -> "--...--."

共有 2 种不同翻译, "--...-." 和 "--...--.".


 */
func uniqueMorseRepresentations(words []string) int {
    morseMap := make(map[string]int)
    for _, word := range words {
        morseStr := ""
        for i := 0; i < len(word); i++ {
            morseStr += morse[word[i]-'a']
        }
        morseMap[morseStr]++
    }
    return len(morseMap)
}
