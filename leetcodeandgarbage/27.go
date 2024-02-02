package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}

func main() {

	nums := []int{0, 1, 2, 2, 3, 2, 4, 0}
	fmt.Println(removeElement(nums, 2))

}

//func removeElement(nums []int, val int) int {
//
//	j := 0
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i] != val {
//			nums[j], nums[i] = nums[i], nums[j]
//			j++
//		}
//	}
//
//	return j
//}
