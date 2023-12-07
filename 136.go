package main

import "fmt"
// Given a non-empty array of integers nums, every element appears twice except for one. 
// Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

// Example 1:

// Input: nums = [2,2,1]
// Output: 1

// Example 2:

// Input: nums = [4,1,2,1,2]
// Output: 4


func singleNumber(nums []int) bool {
	m := make(map[int]int)

	for _, e := range nums {
		m[e]++

	}

	for _, e := range nums {
		if m[e] > len(nums)/2 {
			fmt.Println(e)
			break
		}
	}
	return false
}

func main() {

	//nums := []int{3,3}
	nums := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
	fmt.Println(singleNumber(nums))

}
