func lengthOfLongestSubstring(s string) int {
	Map := make(map[byte]bool)
	Left, Max := 0, 0

	for Right := 0; Right < len(s); Right++ {
		for Map[s[Right]] {
			delete(Map, s[Left])
			Left++
		}
		Map[s[Right]] = true

		//update Max
		temp := Right - Left + 1
		if temp > Max {
			Max = temp
		}
	}
	return Max
}