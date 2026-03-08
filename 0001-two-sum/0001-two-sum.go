func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, num :=  range nums{
        c := target - num

        if idx, ok := m[c]; ok{
            return []int{idx, i}
        }
        m[num] = i
    }
    return []int{}
}