func findDuplicates(nums []int) []int {
	freq := make(map[int]int)
	result := []int{}

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	for key, value := range freq {
		if value > 1 {
			result = append(result, key)
		}
	}

	return result
}