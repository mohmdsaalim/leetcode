func findPeakElement(nums []int) int {
  for i := 1; i<len(nums)-1; i++{
    if nums[i-1] < nums[i] && nums[i] > nums[i+1]{
        return i
    }
  }
    var maxIndex int
    
    for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			maxIndex = i + 1
		}
	}
	return maxIndex
}