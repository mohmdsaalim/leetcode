func mostWordsFound(sentences []string) int {
    maxWords := 0

    for _, s := range sentences{
        wordCount := 1
        for _,ch := range s{
            if ch == ' '{
                wordCount++
            }
        }

        if wordCount > maxWords{
            maxWords = wordCount
        }
    }
    return maxWords
}