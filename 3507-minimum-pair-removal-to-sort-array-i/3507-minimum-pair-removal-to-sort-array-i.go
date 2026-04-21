
func minimumPairRemoval(nums []int) int {
	operations := 0

	for !isNonDecreasing(nums) {
		    minSum := nums[0] + nums[1]
		  index := 0

		for i := 1; i < len(nums)-1; i++ {
		    	sum := nums[i] + nums[i+1]
		 	if sum < minSum {
			    	minSum = sum
				index = i
			}
		}

	    	nums[index] = minSum
		nums = append(nums[:index+1], nums[index+2:]...)
		    operations++
	}
	return operations
}

func    isNonDecreasing(nums []int) bool {
	    for i := 1; i < len(nums); i++ {
		    if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}