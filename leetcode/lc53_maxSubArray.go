package leetcode

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := dp[0]
	for i := 1; i < len(nums); i++ {
		if (nums[i] + dp[i-1]) > nums[i] {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}

	return maxSum

}

func GF_maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

