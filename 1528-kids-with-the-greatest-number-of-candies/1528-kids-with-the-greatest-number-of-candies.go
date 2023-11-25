func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandies := candies[0]
    for i := 1; i < len(candies); i++ {
        if candies[i] > maxCandies {
            maxCandies = candies[i]
        }
    }

    var result []bool
    for i := 0; i < len(candies); i++ {
        if candies[i] + extraCandies >= maxCandies {
            result = append(result, true)
        } else {
            result = append(result, false)
        }
    }
    return result
}