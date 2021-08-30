package acm
//
////一个已经排好序的数组（有重复数字，全是正整数），需要把这个数组切成若干个字数组，
////要求子数组里面的数必须是连续的并且数量至少为3，并且原数组的数全部都要分配到某个子数组里面。问：判断原数组能否切分
//
////  1 1 2 2 3 3  4 4 5   6 6  7
//
//// n 的个数 = n-1 + n+1
//
//// len / 3  剩下的分配
//// 1个  当前是否连续
//// 2个  重复数字拆
//
//func splitArr(arr []int) bool {
//	if len(arr) < 3 {
//		return false
//	}
//	//flag := true
//	cmap := make(map[int]int, len(arr))
//	for i := 0; i < len(arr)-1; i++ {
//		if arr[i+1] - arr[i] == 0 {
//			cmap[arr[i]]++
//		} else {
//			cmap[arr[i]] = 1
//		}
//	}
//	// 按分组取数 map[1]  map[2]
//	group := make(map[int]int, len(arr))
//	g := 0
//	for i := 0; i < len(arr)-2; i++ {
//		if (arr[i+2]-arr[i+1]) == 1 && (arr[i+1]-arr[i]) == 1  {
//
//		}
//	}
//	for _, num := range cmap {
//		if cmap[num] > 0 {
//			for _, g := range group {
//				if group[g] > 0 {
//					cmap[num]--
//				}
//			}
//		}
//	}
//
//
//
//
//	return false
//}
//
//
//

func isPossible(nums []int) bool {
	left := map[int]int{} // 每个数字的剩余个数
	for _, v := range nums {
		left[v]++
	}
	endCnt := map[int]int{} // 以某个数字结尾的连续子序列的个数
	for _, v := range nums {
		if left[v] == 0 {
			continue
		}
		if endCnt[v-1] > 0 { // 若存在以 v-1 结尾的连续子序列，则将 v 加到其末尾
			left[v]--
			endCnt[v-1]--
			endCnt[v]++
		} else if left[v+1] > 0 && left[v+2] > 0 { // 否则，生成一个长度为 3 的连续子序列
			left[v]--
			left[v+1]--
			left[v+2]--
			endCnt[v+2]++
		} else { // 无法生成
			return false
		}
	}
	return true
}

func isTrue(arr []int) bool {
	left := map[int]int{}
	for _, v := range arr {
		left[v]++
	}
	end := map[int]int{}
	for _, v := range arr {
		if left[v] == 0 {
			continue
		}
		if end[v-1] > 0 {
			left[v]--
			end[v-1]--
			end[v]++
		} else if left[v+1] > 0 && left[v+2] > 0 {
			left[v]--
			left[v+1]--
			left[v+2]--
			end[v+2]++
		} else {
			return false
		}
	}
	return true
}
