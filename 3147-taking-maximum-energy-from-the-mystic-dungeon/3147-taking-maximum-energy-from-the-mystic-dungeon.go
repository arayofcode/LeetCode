func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    for i := n-1-k; i >= 0; i-- {
        energy[i] += energy[i+k]
    }

    max := -999999
    for i := n-1; i >= 0; i-- {
        if energy[i] > max {
            max = energy[i]
        }
    }
    return max
}