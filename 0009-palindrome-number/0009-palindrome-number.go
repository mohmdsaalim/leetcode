import "strconv"

func isPalindrome(x int) bool {
    xStr := strconv.Itoa(x)
    // fmt.Println(xStr)
    for idx, jdx := 0, len(xStr)-1; idx <= jdx;{
        if xStr[idx] != xStr[jdx] {return false}
        idx++
        jdx-- 
    }
    return true
}