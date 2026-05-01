func convertToTitle(columnNumber int) string {
    finalStr := ""

    for columnNumber > 0 {
        columnNumber-- 
        remainder := columnNumber % 26
        finalStr = string('A' + remainder) + finalStr
        columnNumber /= 26
    }

    return finalStr
}