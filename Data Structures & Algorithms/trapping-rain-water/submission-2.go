func trap(heights []int) int {
    prefix := make([]int, len(heights))
    curHeight := 0
    for idx, height := range heights {
        prefix[idx] = curHeight 
        curHeight = max(curHeight, height)
    }
    
    suffix := make([]int, len(heights))
    curHeight = 0
    for idx := len(heights)-1; idx >= 0; idx-- {
        height := heights[idx]
        suffix[idx] = curHeight
        curHeight = max(curHeight, height)
    }

    total := 0
    for idx, height := range heights {
       total += max(min(prefix[idx], suffix[idx]) - height, 0)
    }
    return total
}
