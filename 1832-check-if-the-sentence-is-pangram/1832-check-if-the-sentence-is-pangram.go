func checkIfPangram(sentence string) bool {
    a := make([]bool, 26)

    for _, v := range sentence {
        a[v -'a'] = true
    }

    for _, i := range a{
        if !i{
            return false
        }
    }
    return true
}