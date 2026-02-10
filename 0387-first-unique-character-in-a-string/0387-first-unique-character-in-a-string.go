func firstUniqChar(s string) int {
    
    var c = make([]int, 26)

    for i := 0; i < len(s); i++{
        c[s[i] - 'a']++
    }

    for i := 0; i < len(s); i++{
        if c[s[i]-'a'] == 1{
            return i
        }
    }
    return -1
}