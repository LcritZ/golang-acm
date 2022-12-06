package leetcode

import "fmt"


/****
nums := []int{-2,0,3,-5,2,-1}
计算区间和

可以提前算好，后面直接再用区间减区间得到指定区间

 */
type NumArray struct {
    preSums []int
}

func ConstructorSum(nums []int) NumArray {
    numArray := NumArray{
    }
    numArray.preSums = make([]int, len(nums)+1)
    numArray.preSums[0] = 0
    for i := 1; i <= len(nums); i++ {
        numArray.preSums[i] = numArray.preSums[i-1]+nums[i-1]
        fmt.Println(numArray.preSums[i-1])
    }
    fmt.Println(numArray.preSums)
    return numArray
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.preSums[right+1]-this.preSums[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
