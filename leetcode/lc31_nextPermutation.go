package leetcode

/**
排列组合，按照字典顺序，找到下一个，然后

思路：其实就是前后数据的swap，如果当前已经是最大，就需要掉转一下

1,2,3

1,3,2

2,1,3

2,3,1

3,1,2

3,2,1

2,4,3,1

3,1,2,4

[5,4,7,5,3,2]

5,7,2,3,4,5
5,5,2,3,4,7

思路：找到第一个大于当前需要调整的值，再走一遍，从后面找到需要替换的最小值

后面都是降序，然后找到第一个没有保持降序的数，然后用这个数和后面降序对比，找到第一个大于它的数，交换，这样子后面还是降序，但是取了一个刚刚好升级的数过来，然后把他们按升序排列，就变成了升级后的排列

 */
func NextPermutation(nums []int)  {
    if len(nums) == 1 {
        return
    }

    index := -1
    i := len(nums)-2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }

    index = i
    if index >= 0 {
        j := len(nums)-1
        for j >= 0 && nums[index] >= nums[j] {
            j--
        }
        nums[index], nums[j] = nums[j], nums[index]
    }
    reverse31(nums[index+1:])

    return
}

func reverse31(nums []int) {
    for i := 0; i < len(nums)/2; i++ {
        nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
    }
}


func GF_nextPermutation(nums []int) {
    n := len(nums)
    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    if i >= 0 {
        j := n - 1
        for j >= 0 && nums[i] >= nums[j] {
            j--
        }
        nums[i], nums[j] = nums[j], nums[i]
    }
    reverse31(nums[i+1:])
}
