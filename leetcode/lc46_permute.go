package leetcode

/**
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

 */

var permuteResult [][]int
var numsLength int

func Permute(nums []int) [][]int {
	numsLength = len(nums)
	permuteResult = [][]int{}
	if numsLength == 0 {
		return permuteResult
	}
	permuteBackTrace(nums, []int{})
	return permuteResult
}

func permuteBackTrace(nums []int, group []int) {
	if len(group) == numsLength {
		permuteResult = append(permuteResult, group)
		return
	}

	selectMap := make(map[int]struct{}, len(group))
	for i := 0; i < len(group); i++ {
		selectMap[group[i]] = struct{}{}
	}
	for i := 0; i < numsLength; i++ {
		//排除不合条件的
		if _, ok := selectMap[nums[i]]; ok {
			continue
		}
		//用新的group备份存储，避免了使用group需要回溯的情况
		newGroup := append(group, nums[i])
		permuteBackTrace(nums, newGroup)
	}
}

//直接修改了获取数据的数组，这样子就不用过滤，因为数组已经是过滤之后的了
var res [][]int
func Permute2(nums []int) [][]int {
	res = [][]int{}
	permutebackTrack2(nums,len(nums),[]int{})
	return res
}
func permutebackTrack2(nums []int,numsLen int,path []int)  {
	if len(nums)==0{
		p:=make([]int,len(path))
		copy(p,path)
		res = append(res,p)
	}
	for i:=0;i<numsLen;i++{
		//做出选择
		cur:=nums[i]
		path = append(path,cur)
		nums = append(nums[:i],nums[i+1:]...)//直接使用切片
		//下一轮trace
		permutebackTrack2(nums,len(nums),path)
		//撤销选择
		nums = append(nums[:i],append([]int{cur},nums[i:]...)...)//回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]

	}
}


