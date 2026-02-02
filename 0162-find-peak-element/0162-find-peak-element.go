func findPeakElement(nums []int) int {
	max := nums[0]
	maxIdx := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
	}
	return maxIdx
}