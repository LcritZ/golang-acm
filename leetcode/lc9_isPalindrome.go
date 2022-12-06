package leetcode


func isPalindrome(x int) bool {

    if x < 0 || (x != 0 && x%10 == 0) {
        return false
    }

    revertInt := 0
    for x > revertInt {
        revertInt = 10*revertInt+x%10
        x /= 10
    }

    return x == revertInt || x == revertInt/10

}
