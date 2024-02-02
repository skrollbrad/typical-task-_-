package main

import "fmt"

func removeDuplicates(nums []int) int {
	left := 1
	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[right-1] {

			left++
		}
	}

	return left
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 6, 6}
	fmt.Println(removeDuplicates(nums))

}
