package basic

import (
	"fmt"
)

/**
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
如果是则返回true,否则返回false。假设输入的数组的任意两个数字都互不相同。
（ps：我们约定空树不是二叉搜索树）
1 边界条件判断
2 根节点在最后，判断小于，大于分界mid
3 没有比mid大的数情况下，mid初始化是0，或者是第一个大于root，mid也是0，所以要判断小于 root 则递归到 左子树
4 大于的情况下，mid后面的数据都要大于root，否则false
5 符合则对左右子树递归， 注意数组不能越界，根据mid判断是否进行左右递归

 */
func VerifySquenceOfBST( sequence []int ) bool {
	// write code here
	if len(sequence) == 0 {
		return false
	}
	if len(sequence) == 1 {
		return true
	}
	var root int = sequence[len(sequence)-1]
	var mid int
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] > root {
			mid = i
			break
		}
	}
	if mid == 0 && sequence[mid] < root {
		return VerifySquenceOfBST(sequence[:len(sequence)-1])
	}
	for j := mid+1; j < len(sequence)-1; j++ {
		if sequence[j] < root {
			fmt.Println(sequence[j], root)
			return false
		}
	}
	left := true
	right := true
	if mid > 0 {
		left = VerifySquenceOfBST(sequence[:mid])
	}
	if mid < len(sequence) -1 {
		right = VerifySquenceOfBST(sequence[mid:len(sequence)-1])
	}
	fmt.Println(sequence[mid])
	return left && right

}

func VerifySeqenceOfBST2(seq []int) bool {
	if len(seq) == 0 {
		return false
	}
	if len(seq) == 1 {
		return true
	}
	mid := 0
	root := seq[len(seq)-1]
	for i :=0; i < len(seq)-1; i++ {
		if seq[i] > root {
			mid = i
			break
		}
	}
	if mid == 0 && seq[mid] < root {
		return VerifySeqenceOfBST2(seq[:len(seq)-1])
	}
	for i := mid; i < len(seq)-1; i++ {
		if seq[i] < root {
			return false
		}
	}
	left, right := true, true
	if mid > 0 {
		left = VerifySeqenceOfBST2(seq[:mid])
	}
	if mid < len(seq)-1 {
		right = VerifySeqenceOfBST2(seq[mid:len(seq)-1])
	}
	return left && right
}
