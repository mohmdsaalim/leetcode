func backspaceCompare(s string, t string) bool {
    return backspace(s) == backspace(t)
}

func backspace(s string) string {
    ln := 0
    rn := make([]rune, len(s))
    for i := range s {
        if s[i] == '#' {
           ln = max(0, ln-1)
           continue
        }
        rn[ln] = rune(s[i])
        ln++
    }
    return string(rn[:ln])
}