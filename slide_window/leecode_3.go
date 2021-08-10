//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
/*
滑动窗口的典型题目，寻找无重复的最长子串
通过左右指针不断移动
每次都移动右指针，并且不断将字符加入到窗口内
移动过程中，加入窗口前，若发现当前右指针所指的字符已经在窗口中了
那么需要开始移动左指针，不断删除滑动窗口中左指针所指的元素，然后左指针左移
直到右指针所指的元素不再出现在窗口中
此时右指针所指元素可以加入到窗口了，然后继续右移，直到移动到字符串末尾
*/
func lengthOfLongestSubstring(s string) int {
    window := map[byte]struct{}{}
    left := 0
    max := 0
    for right := 0; right < len(s); right++ {
        if _, ok := window[s[right]]; !ok {
            window[s[right]] = struct{}{}
        } else {
            for {
                if _, ok := window[s[right]]; ok {
                    delete(window, s[left])
                    left++
                } else {
                    break
                }
            }
            window[s[right]] = struct{}{}
        }
        if len(window) > max {
            max = len(window)
        }
    }
    return max
}
