func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    matrix := make([][]byte, numRows)
    for i := 0; i < numRows; i++ {
        matrix[i] = []byte{}
    }
    flag := -1
    i := 0
    for j := 0; j < len(s); j++ {
        matrix[i] = append(matrix[i], s[j])
        if i == numRows - 1 || i == 0 {
            flag = -flag
        }
        i += flag
    }
    result := []byte{}
    for i := 0; i < numRows; i++ {
        result = append(result, matrix[i]...)
    }
    return string(result)
}
