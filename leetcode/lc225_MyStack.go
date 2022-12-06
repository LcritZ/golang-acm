package leetcode

type MyStack struct {
    arr1 []int
    arr2 []int
}


func ConstructorStack() MyStack {
    return MyStack{
        arr1: []int{},
        arr2: []int{},
    }
}


func (this *MyStack) Push(x int)  {
    this.arr2 = append(this.arr2, x)
    this.arr2 = append(this.arr2, this.arr1...)
    this.arr1 = []int{}
    this.arr1, this.arr2 = this.arr2, this.arr1
}

func (this *MyStack) Pop() int {
    res := this.arr1[0]
    this.arr1 = this.arr1[1:]
    return res
}


func (this *MyStack) Top() int {
    return this.arr1[0]
}


func (this *MyStack) Empty() bool {
    return len(this.arr1) == 0
}


