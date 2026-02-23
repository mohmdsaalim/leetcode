func arraySign(nums []int) int {
    mult:= 1 
    for _, n := range nums{
        if n > 0{
            n = 1
        }
        if n < 0{
            n = -1
        }
        if n == 0 {
            return 0
        }
        mult *= n 
    }
    return mult
}