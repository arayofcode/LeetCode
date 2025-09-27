func largestTriangleArea(points [][]int) float64 {
    return bruteForce(points)
}

func bruteForce(points [][]int) float64 {
    // Find all triplets and calculate their areas. Return maximum
    // Time: O(n^3)

    maxArea := float64(-1)
    for i := 0; i < len(points) - 2; i++ {
        for j := i+1; j < len(points) - 1; j++ {
            for k := j+1; k < len(points); k++ {
                area := calculateArea(points[i], points[j], points[k])
                if area > maxArea {
                    maxArea = area
                }
            }
        }
    }

    return maxArea
}

func calculateArea(p1 []int, p2 []int, p3 []int) float64 {
    x1, y1 := float64(p1[0]), float64(p1[1])
    x2, y2 := float64(p2[0]), float64(p2[1])
    x3, y3 := float64(p3[0]), float64(p3[1])

    area := (x1 * (y2 - y3) + x2 * (y3 - y1) + x3 * (y1 - y2)) / 2.0
    if area < 0 {
        return -area
    }
    return area
}