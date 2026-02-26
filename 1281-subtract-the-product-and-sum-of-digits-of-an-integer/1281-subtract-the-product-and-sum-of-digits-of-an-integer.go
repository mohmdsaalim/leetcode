func subtractProductAndSum(n int) int {
    m, s := 1, 0
    for n > 0{
        s += n % 10
        m *= n % 10
        n /= 10
    }
    return m - s
}