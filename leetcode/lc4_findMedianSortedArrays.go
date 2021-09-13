package leetcode

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


 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	var _ bool
	//flag := false
	if (len(nums1)+len(nums2))%2 == 0 {
		_ = true
	} else {
		_ = false
	}
	var midIndex int
	midIndex = (len(nums1)+len(nums2))/2
	i := 1
	j:= 1
	for i+j < midIndex && i-1 < len(nums1) && j-1 < len(nums2) {
		if nums1[i-1] <= nums2[j-1] {
			i ++
		} else {
			j++
		}
	}
	//if flag {
	//
	//	if nums1[i-1] > nums2[j-1] {
	//		return float64()
	//	}
	//}
	_ = false

	return 0.0
}
