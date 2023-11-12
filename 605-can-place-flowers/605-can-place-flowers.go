func canPlaceFlowers(flowerbed []int, n int) bool {
    count := 0

    if len(flowerbed) == 1 {
        if flowerbed[0] == 0 {
            count++
        }
    } else if len(flowerbed) == 2 {
        if flowerbed[0] == 0 && flowerbed[1] == 0 {
            count++
        }
    } else {
        for i := 0; i < len(flowerbed) - 1; i++ {
            if i == 0 {
                if flowerbed[0] == 0 && flowerbed[1] == 0 {
                    flowerbed[0] = 1
                    count++
                }
                continue
            }
            if flowerbed[i - 1] == 0 && flowerbed[i] == 0 && flowerbed[i + 1] == 0 {
                flowerbed[i] = 1
                count++
            }
            if i == len(flowerbed) - 2 {
                if flowerbed[i] == 0 && flowerbed[i + 1] == 0 {
                    count++
                }
            }
        }

    }
    return count >= n
}