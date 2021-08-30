package leetcode

/**
最长递增数组
 */
func LengthOfLIS(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	start, end := 0, 0
	for i :=0; i < len(arr); i++ {
		l1, r1 := LISHelper(arr, i, i)
		if r1-l1 > end-start {
			start, end = l1, r1
			i = r1
		}
	}
	return arr[start:end+1]
}

func LISHelper(arr[]int, left, right int) (int, int) {
	for ; right < len(arr) - 1;  {
		if arr[right+1] > arr[right] {
			right = right+1
		} else {
			break
		}
	}
	return left, right
}

func LengthOfLIS2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := 1
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j]+1
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

