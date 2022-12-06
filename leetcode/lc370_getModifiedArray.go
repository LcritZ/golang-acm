package leetcode

func GetModifiedArray(length int, updates [][]int) []int {
    if length == 0 {
        return []int{}
    }

    diff := make([]int, length)
    for i := 0; i < len(updates); i++ {
        x, y, z := updates[i][0], updates[i][1], updates[i][2]
        diff[x] += z
        if y+1 < length {
            diff[y+1] -= z
        }
    }

    nums := make([]int, length)
    nums[0] = diff[0]
    for i := 1; i < length; i++ {
        nums[i] += nums[i-1] + diff[i]
    }
    return nums
}
