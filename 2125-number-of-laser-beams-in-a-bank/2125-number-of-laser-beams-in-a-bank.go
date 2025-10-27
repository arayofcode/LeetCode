func numberOfBeams(bank []string) int {
    count := make([]int, 0, len(bank))
    for _, row := range bank {
        c := 0
        for _, cell := range row {
            if cell == '1' {
                c++
            }
        }
        if c > 0 {
            count = append(count, c)
        }
    }
    beams := 0
    for i := 0; i < len(count) - 1; i++ {
        beams += count[i] * count[i+1]
    }
    return beams
}
