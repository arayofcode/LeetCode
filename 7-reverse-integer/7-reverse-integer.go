func reverse(x int) int {
    var copy int = x
    var result int = 0
    for (copy != 0) {
        result = result * 10 + copy % 10
        copy = copy / 10
    }
    if (result > math.MaxInt32 || result < math.MinInt32) {
        return 0
    }
    return result
}