package leetcode

import "golang-acm/util"

/**
2,3
1,4
1,2,3,4
0,1,2,3

1,3,5,
2,4,6

1,2,3
4,5,6

1,2,4
3,5,6

1,2,3,4,5,6
0,1,2,3


2+3 / 2 = 2.5

2 = index0  3 - index1

nums1  0 1  1-2

nums2 0     1

思路：双指针，小和大向中间靠拢

 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//if len(nums1) == 0 && len(nums2) == 0 {
	//	return 0.0
	//}
	////var _ bool
	//flag := false
	//if (len(nums1)+len(nums2))%2 == 0 {
	//	flag = true
	//} else {
	//	flag = false
	//}
	//var midIndex int
	//midIndex = (len(nums1)+len(nums2))/2
	//s1, e1 := 0, len(nums1)-1
	//s2, e2 := 0, len(nums2)-1
	//
	//for i := 0; i < midIndex; i ++ {
	//
	//}



	//if flag {
	//
	//	if nums1[i-1] > nums2[j-1] {
	//		return float64()
	//	}
	//}
	_ = false

	return 0.0
}

func GF_FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex + 1))
	} else {
		midIndex1, midIndex2 := totalLength/2 - 1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1 + 1) + getKthElement(nums1, nums2, midIndex2 + 1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2 + k - 1]
		}
		if index2 == len(nums2) {
			return nums1[index1 + k - 1]
		}
		if k == 1 {
			return util.Min(nums1[index1], nums2[index2])
		}
		half := k/2
		newIndex1 := util.Min(index1 + half, len(nums1)) - 1
		newIndex2 := util.Min(index2 + half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}


