func findWordsContaining(words []string, x byte) []int {
    res := []int{}
    for idx, word := range words {
        for _, w := range word {
            if string(w) == string(x) {
                res = append(res, idx)
                break
            }
        }
    }
    return res
}