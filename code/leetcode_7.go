func reverse(x int) int {
    if x == 0 {
        return 0
    }
    negative := false
    if x < 0 {
        negative = true
        x = -x
    }
    num := 0
    sum := 0
    for x > 0 {
        num = x % 10
        sum = sum * 10 + num
        x /= 10
    }
    if negative {
        if sum > 1 << 31 {
            return 0
        }
        return -sum
    }
    if sum > 1 << 31 - 1 {
        return 0
    }
    return sum
}
