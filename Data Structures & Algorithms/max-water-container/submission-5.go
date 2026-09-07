func maxArea(heights []int) int {
    maxArea := 0
    for i, j := 0, len(heights)-1; i != j; {
        height := min(heights[i], heights[j])
        base := j - i
        area := height * base
        if area > maxArea {
            maxArea = area
        }
        if heights[i] < heights[j] {
            i++
        } else {
            j--
        }
    }
    return maxArea
}
