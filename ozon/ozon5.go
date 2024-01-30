package main

import (
	"fmt"
)

func work(nums []int, count int) {
	newArray := []int{}
	var lowCount, upperCount int
	is := true
	for i := 0; i < count; i++ {
		//newArray = append(newArray, newArray[:1]...)
		for is {
			if nums[i+1]-nums[i] == -1 {
				lowCount++
				if nums[i+2]-nums[i+1] != -1 {
					newArray = append(newArray, newArray[:1]...)
					newArray = append(newArray, lowCount*-1)
					nums = append(nums[:lowCount], nums[lowCount+1:]...)
				}

			}
			if nums[i+1]-nums[i] == 1 {
				upperCount++
				if nums[i+2]-nums[i+1] != 1 {
					newArray = append(newArray, newArray[:1]...)
					newArray = append(newArray, upperCount)
					nums = append(nums[:upperCount], nums[upperCount+1:]...)
				}
			}

		}

	}
}
func main() {

	var numOfSets, count, num int
	slice := []int{}
	fmt.Println("Введите количество наборов входных данных")
	fmt.Scan(&numOfSets)

	for i := 1; i <= numOfSets; i++ {
		fmt.Println("Введите количество элементов")
		fmt.Scan(&count)
		for j := 0; j < count; j++ {

			fmt.Scan(&num)

			slice = append(slice, num)

		}
		work(slice, count)
	}
}
