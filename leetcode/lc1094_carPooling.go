package leetcode

/**
假设你是一位顺风车司机，车上最初有capacity个空座位可以用来载客。由于道路的限制，车只能向一个方向行驶（也就是说，不允许掉头或改变方向，你可以将其想象为一个向量）。

这儿有一份乘客行程计划表trips[][]，其中trips[i] = [num_passengers, start_location, end_location]包含了第 i 组乘客的行程信息：

必须接送的乘客数量；
乘客的上车地点；
以及乘客的下车地点。
这些给出的地点位置是从你的初始出发位置向前行驶到这些地点所需的距离（它们一定在你的行驶方向上）。

请你根据给出的行程计划表和车子的座位数，来判断你的车是否可以顺利完成接送所有乘客的任务（当且仅当你可以在所有给定的行程中接送所有乘客时，返回true，否则请返回 false）

输入：trips = [[2,1,5],[3,3,7]], capacity = 4
输出：false

注意：边界条件，也就是上车下车是0开始还是1开始，全部站台是多少

 */
func CarPooling(trips [][]int, capacity int) bool {
    temp := make([]int, 1000)
    for i := 0; i < len(trips); i++ {
        x, y, z := trips[i][0], trips[i][1], trips[i][2]
        temp[y] += x
        temp[z] -= x
    }
    curr := 0
    for i := 0; i < len(temp); i++ {
        curr += temp[i]
        if curr > capacity {
            return false
        }
    }
    if curr > 0 {
        return false
    }

    return true
}
