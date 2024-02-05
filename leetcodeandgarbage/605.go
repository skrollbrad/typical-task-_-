package main

// Time complexity:  O(n)
// Space complexity: O(1)
func canPlaceFlowers(flowerbed []int, n int) bool {
	// remained_plots := 0

	i := 0

	for i < len(flowerbed) && n > 0 {
		if i == 0 {
			if flowerbed[i] == 0 && len(flowerbed) == 1 {
				flowerbed[i] = 1
				n--
			} else if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else if i == len(flowerbed)-1 {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}

		i++

	}

	if n == 0 {
		return true
	}
	return false

}
func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	canPlaceFlowers(flowerbed, n)
}
