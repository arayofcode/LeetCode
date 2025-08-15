import "math"

func isPowerOfFour(n int) bool {
    x := math.Ceil(math.Log2(float64(n)) / 2)
    return (n >= 4 || n == 1) && (float64(n) == math.Pow(4, x))
}
