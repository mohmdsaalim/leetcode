func isPerfectSquare(num int) bool {
    start := 1
    end := num

    for end >= start {
        mid := end/2 + start/2 + (start%2 + end%2)/2
        if mid == num/mid && num % mid == 0 {
            return true
        } else if mid > num/mid {
            end = mid - 1
        } else {
            start = mid + 1
        }
    }
    return false
}