func judgeCircle(moves string) bool {
    c,b := 0, 0
    for _, ccb := range moves {
        switch ccb {
            case 'R':
            c++
             case 'L':
            c--
            case 'U':
            b++
             case 'D':
            b --
        }
    }
    return c == 0 && b == 0
}