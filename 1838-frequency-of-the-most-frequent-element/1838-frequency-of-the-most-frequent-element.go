import "sort"

func maxFrequency(nums []int, k int) int {

    sort.Ints(nums)

    l := 0
    sum := 0
    res := 1

    for r := 0; r < len(nums); r++ {

        sum += nums[r]

        for (nums[r]*(r-l+1) - sum) > k {
            sum -= nums[l]
            l++
        }

        res = max(res, r-l+1)
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}