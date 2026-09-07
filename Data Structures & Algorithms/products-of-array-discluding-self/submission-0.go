func productExceptSelf(nums []int) []int {
    preProd := make([]int, len(nums))
    preProd[0] = 1
    for idx := 1; idx < len(nums); idx++ {
        preProd[idx] = preProd[idx-1] * nums[idx-1]
    }
    postProd := make([]int, len(nums))
    postProd[len(nums)-1] = 1
    for idx := len(nums)-2; idx >= 0; idx-- {
        postProd[idx] = postProd[idx+1] * nums[idx+1]
    }
    result := make([]int, len(nums))
    for idx := range nums {
        result[idx] = preProd[idx] * postProd[idx]
    }
    return result
}
