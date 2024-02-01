package main

import "fmt"

func main() {
	scoores := []int{25, 23, 58, 19, 77, 54}
	length := len(scoores)
	fmt.Println(minValue(scoores, length))
}
func minValue(scores []int, lenght int) []int {
	var minVal = 0
	// 50 40 30 60 70

	for i := 0; i < lenght; i++ {
		minVal = i
		for j := i; j < lenght; j++ {
			if scores[minVal] > scores[j] {
				scores[minVal], scores[j] = scores[j], scores[minVal]
			}
		}

	}
	return scores

}
