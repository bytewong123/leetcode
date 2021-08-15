func intToRoman(x int) string {
    type tuple struct {
        num int
        roman string
    }
    romanToInt := []tuple{
        {num : 1000, roman : "M"},
        {num : 900, roman : "CM"},
        {num : 500, roman : "D"},
        {num : 400, roman : "CD"},
        {num : 100, roman : "C"},
        {num : 90, roman : "XC"},
        {num : 50, roman : "L"},
        {num : 40, roman : "XL"},
        {num : 10, roman : "X"},
        {num : 9, roman : "IX"},
        {num : 5, roman : "V"},
        {num : 4, roman : "IV"},
        {num : 1, roman : "I"},
    }
    result := ""
    for x > 0 {
        for i := 0; i < len(romanToInt); i++ {
            if romanToInt[i].num > x {
                continue
            }
            x -= romanToInt[i].num
            result += romanToInt[i].roman
            break
        }
    }
    return result
}
