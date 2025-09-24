func sumZero(n int) []int {
    results := []int{}
    if n % 2 == 1 {
        results = append(results, 0)
    }
    for i := 1; i <= n / 2; i++ {
        results = append(results, i, i * -1)
    }
    return results
}