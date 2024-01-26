package main

import "fmt"

// Given an integer n, return true if it is a power of two. Otherwise, return false.

// An integer n is a power of two, if there exists an integer x such that n == 2x.
// Example 1:

// Input: n = 1
// Output: true
// Explanation: 20 = 1

// Example 2:

// Input: n = 16
// Output: true
// Explanation: 24 = 16

// Example 3:

// Input: n = 3
// Output: false


func isPowerOfTwo(n int) bool {
	//ok := false
	if n == 0 {
		return false
	}
		var sum int
		for n != 1 {
			if n%2 == 0 {
				n = n / 2
				sum++
			} else {
				return false
			}

		}
	
	return true
}
// через факториал( не моё)

// func isPowerOfTwo(n int) bool {
//     if n <= 0 {
//         return false
//     }
//     return n&(n-1) == 0
// }

func main() {

	fmt.Println(isPowerOfTwo(64))

}
