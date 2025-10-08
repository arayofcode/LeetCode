func successfulPairs(spells []int, potions []int, success int64) []int {
    return binarySearch(spells, potions, success)
}

func bruteForce(spells []int, potions []int, success int64) []int {
    results := make([]int, len(spells))
    for i, spell := range spells {
        count := 0
        for _, potion := range potions {
            if int64(spell) * int64(potion) >= success {
                count++
            }
        }
        results[i] = count
    }
    return results
}

func binarySearch(spells []int, potions []int, success int64) []int {
    slices.Sort(potions)
    n := len(potions)
    results := make([]int, len(spells))
    for i, spell := range spells {
        minPotion := int(success / int64(spell))
        if success % int64(spell) > 0 {
            minPotion += 1
        }
        idx, _ := slices.BinarySearch(potions, minPotion)
        if idx >= n {
            results[i] = 0
        } else {
            results[i] = n-idx
        }
    }
    return results
}