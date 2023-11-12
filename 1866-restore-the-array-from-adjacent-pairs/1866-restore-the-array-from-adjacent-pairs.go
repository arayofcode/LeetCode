func restoreArray(adjacentPairs [][]int) (results []int) {
    // Approach 1: Put everything in map, find key with only one element
    // Start from here, do DFS, append to result as you discover new elements
    var elementsMap = map[int][]int{}
    for _, pair := range adjacentPairs {
        a, b := pair[0], pair[1]
        elementsMap[a] = append(elementsMap[a], b)
        elementsMap[b] = append(elementsMap[b], a)
    }
    for key, adjacent := range elementsMap {
        if len(adjacent) == 1 {
            results = append(results, key)
            results = append(results, adjacent[0])
            break
        }
    }
    i := results[1]
    prev := results[0]
    for {
        pair := elementsMap[i]
        if len(pair) == 1 {
            break
        } else {
            a, b := pair[0], pair[1]
            if a == prev {
                results = append(results, b)
                prev = i
                i = b
            } else {
                results = append(results, a)
                prev = i
                i = a
            }
        }
    }
    return
}

