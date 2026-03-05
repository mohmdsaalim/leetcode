func triangleType(nums []int) string {
    eq := 0
    if nums[0] == nums[1] || nums[0] == nums[2] {
        eq++
    }
    if nums[1] == nums[2] {
        eq++
    }
    if nums[0] + nums[1] <= nums[2] ||
    nums[0] + nums[2] <= nums[1] ||
    nums[1] + nums[2] <= nums[0] {
        return "none"
    }
    if eq == 0 {
        return "scalene"
    }
    if eq == 1 {
        return "isosceles"
    }
    return "equilateral"
}