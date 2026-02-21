func differenceOfSum(nums []int) int {
    sum, sumdig := 0, 0 
    for _, n := range nums{
        sum += n 
        for n > 0{
            sumdig += n % 10
        n /= 10
        }
    }
    res := sum - sumdig
    if res< 0 {
        return -res
    }
    return res
}