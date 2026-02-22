func differenceOfSums(n int, m int) int {
    ts := n * (n + 1) / 2
    dc := n / m
    ds := m * dc * (dc + 1) / 2
    return ts - 2 * ds
}