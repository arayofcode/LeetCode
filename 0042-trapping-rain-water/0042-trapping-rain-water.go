func trap(height []int) int {
    left := 0
    right := len(height) - 1
    leftMax, rightMax := 0, 0

    result := 0
    for left < right {
        if height[left] <= height[right] {
            leftMax = max(leftMax, height[left])
            result += leftMax - height[left]
            left++
        } else {
            rightMax = max(rightMax, height[right])
            result += rightMax - height[right]
            right--
        }
    }
    return result
}

