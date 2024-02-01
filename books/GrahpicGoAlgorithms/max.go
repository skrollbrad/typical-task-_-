package main

import "fmt"

func main() {
	scores := []int{14, 12, 55, 34, 96, 83, 23}
	fmt.Println(max(scores))
}

func max(scores []int) []int {

	for i := 0; i < len(scores); i++ {
		for j := i; j < len(scores); j++ {
			if scores[i] > scores[j] {
				scores[j], scores[i] = scores[i], scores[j]

			}

		}
	}
	fmt.Println(scores)
	return scores

}
