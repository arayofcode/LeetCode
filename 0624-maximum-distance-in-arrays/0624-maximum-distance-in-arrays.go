func maxDistance(arrays [][]int) int {
    currMin, currMax := arrays[0][0], arrays[0][len(arrays[0]) - 1]
    dist := -1
    for i := 1; i < len(arrays); i++ {
        array := arrays[i]
        dist = max(dist, array[len(array) - 1] - currMin, currMax - array[0])
        currMin = min(currMin, array[0])
        currMax = max(currMax, array[len(array) - 1])
    }
    return dist
}