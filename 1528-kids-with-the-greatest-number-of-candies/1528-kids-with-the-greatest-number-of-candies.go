func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandies := candies[0]
    length := len(candies)
    for i := 1; i < length; i++ {
        if candies[i] > maxCandies {
            maxCandies = candies[i]
        }
    }

    var results = make([]bool, len(candies))
    for i := 0; i < length; i++ {
        if candies[i] + extraCandies >= maxCandies {
            results[i] = true
        }
    }
    return results
}