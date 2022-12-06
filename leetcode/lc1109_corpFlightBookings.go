package leetcode

/**
这里有n个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表bookings ，表中第i条预订记录bookings[i] = [firsti, lasti, seatsi]意味着在从 firsti到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi个座位。

请你返回一个长度为 n 的数组answer，里面的元素是每个航班预定的座位总数。

输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]

思路：差分数组，确定好差分数组，再推演

 */
func CorpFlightBookings(bookings [][]int, n int) []int {
    res := make([]int, n)
    temp := make([]int, n)

    for i := 0; i < len(bookings); i++ {
        booking := bookings[i]
        x,y,z := booking[0], booking[1], booking[2]
        if x >= 1 {
            temp[x-1] += z
        }
        if y < n {
            temp[y] -= z
        }
    }

    res[0] = temp[0]
    for i := 1; i < n; i++ {
        res[i] = res[i-1]+temp[i]
    }
    return res
}

