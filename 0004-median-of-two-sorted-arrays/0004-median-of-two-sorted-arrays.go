func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    return mergeMedian(nums1, nums2)
}

// Approach 1 - Merge and find middle element
// Merge the two arrays
// Median is middle element, or average of two in the middle
func mergeMedian(nums1 []int, nums2 []int) float64 {
    merged := mergeSlices(nums1, nums2)

    if len(merged) % 2.0 != 0.0 {
        return float64(merged[len(merged) / 2])
    }
    l := len(merged)
    return (float64(merged[l / 2]) + float64(merged[l / 2 - 1]))/2.0
}

func mergeSlices(nums1 []int, nums2[]int) []int {
    var result []int
    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            result = append(result, nums1[i])
            i++
        } else {
            result = append(result, nums2[j])
            j++
        }
    }
    if i < len(nums1) {
        result = append(result, nums1[i:]...)
    }
    if j < len(nums2) {
        result = append(result, nums2[j:]...)
    }
    return result
}