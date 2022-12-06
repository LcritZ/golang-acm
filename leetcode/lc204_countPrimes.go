package leetcode

func CountPrimes(n int) int {
    nums := []int{}
    numMap := make([]bool, n)
    for i := range numMap {
        numMap[i] = true
    }

    for i := 2; i < n; i++ {
        if numMap[i] {
            nums = append(nums, i)
        }
        for _, j := range nums {
            if i*j >= n {
                break
            }
            numMap[i*j] = false
            if i%j == 0 {
                break
            }
        }
    }

    return len(nums)
}

//需要修改
//func countPrimes(n int) int {
//    if n <= 2 {
//        return 0
//    }
//
//    if n <= 3 {
//        return 1
//    }
//
//    nums := []int{2,3}
//    for i := 4; i < n; i++ {
//        if nums[len(nums)-1]*nums[len(nums)-2] > n {
//            fmt.Println(nums)
//            return len(nums)
//        }
//        flag := true
//        for _, x := range nums {
//            if i % x == 0  {
//                flag = false
//                break
//            }
//        }
//        if flag {
//            nums = append(nums, i)
//        }
//    }
//
//    return len(nums)
//}
