package main

import "fmt"

func removeElement(nums []int, val int) int {
	var counter int

	for i := 0; i < len(nums); i++ {

		for j := i; j < len(nums); j++ {
			if nums[i] == val {
				nums[j], nums[i] = nums[i], nums[j]

			}

		}
	}
	for _, v := range nums {
		if v == val {
			counter++
		}
	}

	nums = append(nums[:0], nums[:len(nums)-counter]...)

	fmt.Println(nums)
	return len(nums)
	//fmt.Println(counter)

}

func main() {

	nums := []int{0, 1, 2, 2, 3, 2, 4, 0}
	fmt.Println(removeElement(nums, 2))

}
