import "strings"
func countKeyChanges(s string) int {
    rn := []rune(strings.ToLower(s))
    count := 0
    for i := 0; i < len(rn)-1; i++ {
        if rn[i] != rn[i+1] {
            count++
        }
    }
    return count
}