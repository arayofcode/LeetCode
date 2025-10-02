func maxBottlesDrunk(numBottles int, numExchange int) int {
    result := 0
    empty := 0
    for numBottles + empty >= 0 {
        result += numBottles
        empty += numBottles
        numBottles = 0
        if empty >= numExchange {
            empty -= numExchange
            numBottles++
            numExchange++
        } else {
            break
        }
    }
    return result
}