func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    origin := x
    newX := 0
    base := 10
    for x > 0 {
        extra := x % 10
        newX = extra + base * newX
        x /= 10
    }
    return newX == origin
}
