func maxArea(height []int) int {
    if len(height) < 2 {
        return 0
    }
    left, right := 0, len(height) - 1
    maxVal := 0
    for left < right {
        area := (right - left) * min(height[left], height[right])
        if area > maxVal {
            maxVal = area
        }
        if height[left] > height[right] {
            right--
        } else {
            left++
        }
    }
    return maxVal
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
