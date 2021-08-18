func isValid(s string) bool {
    stack := make([]byte, 0, len(s))
    for i := 0; i < len(s); i++ {
        if len(stack) == 0 {
            stack = append(stack, s[i])
            continue
        }
        if (s[i] == ')' && stack[len(stack) - 1] == '(') ||
        (s[i] == ']' && stack[len(stack) - 1] == '[') ||
        (s[i] == '}' && stack[len(stack) - 1] == '{') {
            stack = stack[:len(stack) - 1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack) == 0
}
