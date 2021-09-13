package basic

// 小顶堆法,找top k
func TopKByMinHeap(nums []int, k int) []int {
	length := len(nums)
	// 数组长度小于k,直接返回
	if length < k {
		return nums
	}

	// 数组前k个数据取出,并生成小顶堆
	minHeap := make([]int, 0)
	minHeap = append(minHeap, nums[:k]...)
	CreateMinHeap(minHeap)

	// 遍历数组剩余数据,大于堆顶数据时,替换堆顶,重新维护小顶堆
	for i := k; i < length; i++ {
		if nums[i] > minHeap[0] {
			minHeap[0] = nums[i]
			MinHeapify(minHeap, 0, k)
		}
	}

	return minHeap
}

// 自底向上创建小顶堆
func CreateMinHeap(nums []int) {
	length := len(nums)
	for i := length - 1; i >= 0; i-- {
		MinHeapify(nums, i, length)
	}
}

// 维护小顶堆
func MinHeapify(nums []int, posIndex, length int) {
	// 堆左孩子节点index
	leftIndex := 2*posIndex + 1
	// 堆右孩子节点index
	rightIndex := 2*posIndex + 2
	// 当前节点以及左右孩子节点中最小值的index, 初始化为当前节点index
	minIndex := posIndex
	// 左孩子存在并且小于当前节点值时, 最小值index替换为左孩子index
	if leftIndex < length && nums[leftIndex] < nums[minIndex] {
		minIndex = leftIndex
	}
	// 右孩子存在并且小于当前节点值时, 最小值index替换为右孩子index
	if rightIndex < length && nums[rightIndex] < nums[minIndex] {
		minIndex = rightIndex
	}
	// 最小值节点index不等于当前节点时,替换当前节点和其中较小孩子节点值
	if minIndex != posIndex {
		nums[posIndex], nums[minIndex] = nums[minIndex], nums[posIndex]
		MinHeapify(nums, minIndex, length)
	}
}

