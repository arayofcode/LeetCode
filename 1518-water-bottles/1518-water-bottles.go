func numWaterBottles(numBottles int, numExchange int) int {
    drinkable := numBottles
    undrinkable := 0
    result := 0
    for drinkable > 0 {
        result += drinkable
        undrinkable += drinkable
        drinkable = undrinkable / numExchange 
        undrinkable -= drinkable * numExchange
    }
    return result
}