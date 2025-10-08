func successfulPairs(spells []int, potions []int, success int64) []int {
    return replaceSpells(spells, potions, success)
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
        minPotion := int((success + int64(spell) - 1) / int64(spell))
        idx, _ := slices.BinarySearch(potions, minPotion)
        results[i] = n-idx
    }
    return results
}

func replaceSpells(spells []int, potions []int, success int64) []int {
    slices.Sort(potions)
    n := len(potions)
    for i, spell := range spells {
        minPotion := int(success / int64(spell))
        if success % int64(spell) > 0 {
            minPotion += 1
        }
        idx, _ := slices.BinarySearch(potions, minPotion)
        if idx >= n {
            spells[i] = 0
        } else {
            spells[i] = n-idx
        }
    }
    return spells
}