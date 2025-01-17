// There can be multiple duplicates. Use a hash table to store
// last index. Maintain overall minimum cards found so far
// When duplicate is found, compare with existing min
func minimumCardPickup(cards []int) int {
    minCards := -1
    cardIndex := make(map[int]int)
    for i, card := range cards {
        if index, found := cardIndex[card]; found {
            if minCards == -1 {
                minCards = i - index + 1
            } else {
                minCards = min(minCards, i - index + 1)
            }
        }
        cardIndex[card] = i
    }
    return minCards
}