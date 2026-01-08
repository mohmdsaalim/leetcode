func maximumWealth(accounts [][]int) int {
    maxWealth := 0

    for _,c := range accounts{
        sum := 0
    
    for _,m := range c{
        sum+=m
    }
    if sum > maxWealth{
        maxWealth = sum
    }
    }
    return maxWealth
}