package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// Используем массив left, чтобы хранить произведение всех элементов слева от данного индекса
	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	// Используем переменную rightProduct, чтобы хранить произведение всех элементов справа от данного индекса
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] = left[i] * rightProduct
		rightProduct *= nums[i]
	}

	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result) // Выведет: [24 12 8 6]
}
