func canPlaceFlowers(flowerbed []int, n int) bool {
    for i := 0; i < len(flowerbed); i++ {
        prevZero := i == 0 || flowerbed[i-1] == 0
        nextZero := i == len(flowerbed) - 1 || flowerbed[i+1] == 0
        if flowerbed[i] == 0 && prevZero && nextZero {
            n--
            i++
        }
    }
    return n <= 0
}
