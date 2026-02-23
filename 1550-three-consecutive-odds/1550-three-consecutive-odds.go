func threeConsecutiveOdds(arr []int) bool {
    var w int

    for i := 0; i < 3 && i < len(arr); i++{
        w += arr[i] % 2
    }
    if w == 3 {
        return true
    }
    for i := 3; i < len(arr); i++{
        w += arr[i] % 2 
        w -= arr[i-3] % 2 

        if w == 3 {
            return true
        }
    }
    return false
}