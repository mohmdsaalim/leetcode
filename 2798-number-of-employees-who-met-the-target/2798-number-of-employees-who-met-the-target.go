func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    var employeesCount int

    for _, hour := range hours {
        if hour >= target {
            employeesCount++
        }
    }

    return employeesCount
}