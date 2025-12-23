func twoSum(nums []int, target int) []int {
    newarray := []int{}
    for i := 0; i < len(nums);i++{
        for j := i+1; j < len(nums); j++{
                if nums[i] + nums[j] == target{
                    newarray = append(newarray, i,j)
                }
        }
    }
    return newarray
    }
