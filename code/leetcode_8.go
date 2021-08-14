func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if len(s) == 0 {
        return 0
    }
    negative := false
    nums := []byte{}
    for i := 0; i < len(s); i++ {
        if i == 0 {
            if s[i] == '-' {
                negative = true
                continue
            }
            if s[i] == '+' {
                continue
            }
        }
        if s[i] >= '0' && s[i] <= '9' {
            nums = append(nums, s[i])
        } else {
            break
        }
    }
    return strToInt(negative ,nums)
}

func strToInt(negative bool, nums []byte) int {
    idx := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == '0' {
            idx++
        } else {
            break
        }
    }
    nums = nums[idx:]
    num := 0
    base := 1
    for i := len(nums) - 1; i >= 0; i-- {
        if notInArea(base, negative) {
            return limit(negative)
        } 
        num += base * charToInt(nums[i])
        if notInArea(num, negative) {
            return limit(negative)
        }
        base *= 10
    }
    if negative {
        return -num
    }
    return num
}

func notInArea(num int, negative bool) bool {
    return (negative && num >= 1 << 31) || (!negative && num >= 1 << 31 - 1)
}

func limit(negative bool) int {
    if negative {
        return -(1 << 31)
    }
    return 1 << 31 - 1
}

func charToInt(c byte) int {
    if c == '0' {
        return 0
    }
    if c == '1' {
        return 1
    }
    if c == '2' {
        return 2
    }
    if c == '3' {
        return 3
    }
    if c == '4' {
        return 4
    }
    if c == '5' {
        return 5
    }
    if c == '6' {
        return 6
    }
    if c == '7' {
        return 7
    }
    if c == '8' {
        return 8
    }
    if c == '9' {
        return 9
    }
    return -1
}
