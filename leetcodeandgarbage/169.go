package main

import "fmt"

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times.
//  You may assume that the majority element always exists in the array.

// Example 1:

// Input: nums = [3,2,3]
// Output: 3

// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, e := range nums {
		m[e]++

	}
	var s int
	for _, e := range nums {
		if m[e] > len(nums)/2 {
			s = e
			break
		}
	}
	return s
}
func main() {

	//nums := []int{3,3}
	nums := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
	fmt.Println(containsDuplicate(nums))

}
