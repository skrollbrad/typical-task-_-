package main

import (
	"fmt"
	"sort"
)

// 1.сортируем массив, тем самым, мы можем сравнить current и next. Если разница больше 1, значит там что-то пропущено
// далее, если пропущено, вычитаем из next 1 или прибавляем в current 1 до тех пор, пока next -1 || current + 1 != number
// 2. если бы были значение от 1 до 10, то можно было бы завести мапу и посчитать для каждого числа количество повторений m[elem]++
// соответственно, если у числа, допустим 8 m[elem] == 0 , то добавить туда значение этого key.
func missingNumber(nums []int) int {

	sort.Ints(nums)
	if len(nums) > 0 && nums[0] != 0 {
		return 0
	}
	for i, j := 0, 1; i < len(nums); {

		if j == len(nums) || nums[j]-nums[i] > 1 {
			return nums[i] + 1
		}
		i++
		j++
	}

	return 0
}
func main() {

	arrays := []int{3, 0, 1}

	fmt.Println(missingNumber(arrays))
}
