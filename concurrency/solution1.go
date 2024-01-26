package main

import (
	"fmt"
	"time"
)

func main() {

	go fmt.Println(MaxSum([]int{1, 2, 3}, []int{10, 20, 50}))
	go fmt.Println(MaxSum([]int{3, 2, 1}, []int{1, 2, 3}))
	time.Sleep(time.Millisecond)
}

func MaxSum(nums1, nums2 []int) []int {
	s1, s2 := 0, 0

	go sumParallel(nums1, &s1)
	go sumParallel(nums2, &s2)

	time.Sleep(time.Millisecond)
	if s1 >= s2 {
		return nums1
	}
	return nums2

}
func sumParallel(nums []int, res *int) {
	*res = sum(nums)
}
func sum(nums []int) int {
	sum := 0

	for _, v := range nums {
		sum += v

	}

	return sum

}
