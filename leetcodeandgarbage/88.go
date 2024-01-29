package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	for n != 0 {
		nums1[m] = nums2[n-1]
		m++
		n--
	}
	sort.Ints(nums1)
	fmt.Println(nums1)
}

func main() {
	var m, n int
	m = 3
	n = 3
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, m, nums2, n)

}
