// arr[i] = arr[i] ^ arr[i+i], hence arr[i] is present twice
// derived is valid if all elements XOR-ed results in 0
func doesValidArrayExist(derived []int) bool {
    // Slightly less number of instructions
    n := len(derived)
    for i := 1; i < n; i++ {
        derived[i] ^= derived[i-1]
    }
    return derived[n - 1] == 0
}