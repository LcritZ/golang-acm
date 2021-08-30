package acm

func BinarySearch(target int, array []int) bool {
	midIndex := len(array)/2
	if len(array) == 0 {
		return false
	}
	if len(array) == 1 {
		return target == array[0]
	}
	if target < array[midIndex] {
		return BinarySearch(target, array[0:midIndex])
	}
	if target == array[midIndex] {
		return true
	}
	if target > array[midIndex] {
		return BinarySearch(target, array[midIndex:len(array)])
	}
	return false
}

func Find( target int ,  array [][]int ) bool {
	if len(array) == 0 {
		return false
	}
	//普通二分查找 对于array[n][m], 复杂度m*log2n
	//for i := 0; i < len(array); i++ {
	//	findRes := BinarySearch(target, array[i])
	//	if findRes {
	//		return true
	//	}
	//}

	//更优解：复杂度n+m
	if len(array[0]) == 1 {
		return BinarySearch(target, array[0])
	}
	for i := 0; i < len(array); i++ {
		for j := len(array[0])-1; j >= 0; j-- {
			value := array[i][j]
			if target < value {
				continue
			}
			if target == value {
				return true
			}
			if target < value {
				break
			}
		}
	}

	return false
}

