func canAliceWin(nums []int) bool {
    var a,b int
    for _, v := range nums{
        if v < 10 {
            a +=v
        }else{
            b+=v
        }
    }
    return a != b
}