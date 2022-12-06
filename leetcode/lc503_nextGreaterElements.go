package leetcode

func NextGreaterElements(nums []int) []int {
    length := len(nums)
    res := make([]int, len(nums))
    stack := []int{}
    for i := 2*length - 1; i >= 0; i-- {
        num := nums[i%length]
        for len(stack) > 0 && num >= stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) > 0 {
            res[i%length] = stack[len(stack)-1]
        } else {
            res[i%length] = -1
        }
        stack = append(stack, num)
    }
    return res
}
