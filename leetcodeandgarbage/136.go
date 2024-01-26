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

func singleNumber(nums []int) int {
	//m := make(map[int]int)
	// p.s пытался всё в одном цикле сделать, засиделся, в итоге гугланул пример на с++ и сразу понял ошибку
	var count int
	for _, e := range nums {
		count = count | e
		
	

	}

	// for key, e := range m {
	// 	if e == 1 {
	// 		return key
	// 	}
	// }
	return count
}

func main() {

	//nums := []int{3,3}
	nums := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
	fmt.Println(singleNumber(nums))

}
