func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    left, right := 0, 0
    for right < len(nums) {
        if nums[left] == nums[right] {
            right++
            continue
        } 
        left++
        nums[left] = nums[right]
        right++
    }
    return left + 1
}
