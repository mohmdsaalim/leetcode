func findPeakElement(nums []int) int {
    var maxIndex int

    for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			maxIndex = i + 1
		}
	}
	return maxIndex
}