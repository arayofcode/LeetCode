func maxArea(height []int) int {
    area := 0
    left, right := 0, len(height) - 1
    for left < right {
        curArea := (right - left) * min(height[left], height[right])
        if curArea > area {
            area = curArea
        }
        if height[left] > height[right] {
            right--
        } else {
            left++
        }
    }
    return area
}

// // Approach 1 - Brute Force:
// // Look at all possible line pairs and find the maximum area
// // Time Complexity: O(n^2), Space: O(1)
// func bruteForce(height []int) int {
//     area := 0
//     for i := 0; i < len(height); i++ {
//         for j := i+1; j < len(height); j++ {
//             currArea := calculateArea(height, i, j)
//             if currArea > area {
//                 area = currArea
//             }
//         }
//     }
//     return area
// }

// // Approach 2 - Two Pointers:
// // Goal is to maximize the distance and heights
// // Start with leftmost and rightmost lines
// // Switch the smaller one, and find max area
// // Time: O(n), Space: O(1)
// func twoPointers(height []int) int {
//     area, left, right := 0, 0, len(height) - 1
//     for left < right {
//         currArea := calculateArea(height, left, right)
//         if currArea > area {
//             area = currArea
//         }
//         if height[left] < height[right] {
//             left++
//         } else {
//             right--
//         }
//     }
//     return area
// }

// func calculateArea(height []int, left int, right int) int {
//     if height[left] < height[right] {
//         return (right - left) * height[left]
//     }
//     return (right - left) * height[right]
// }