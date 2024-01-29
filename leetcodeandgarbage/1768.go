package main

import "fmt"

// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting
// with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.
// Return the merged string.
//
// Example 1:
//
// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r
// Example 2:
//
// Input: word1 = "ab", word2 = "pqrs"
// Output: "apbqrs"
// Explanation: Notice that as word2 is longer, "rs" is appended to the end.
// word1:  a   b
// word2:    p   q   r   s
// merged: a p b q   r   s
func mergeAlternately(word1 string, word2 string) string {
	word1slice := []int32{}
	word2slice := []int32{}
	allwordsslice := []int32{}
	for _, v := range word1 {
		word1slice = append(word1slice, v)

	}
	for _, v := range word2 {
		word2slice = append(word2slice, v)

	}

	for i := 0; i < len(word1); i++ {

		allwordsslice = append(allwordsslice, word1slice[i])
		allwordsslice = append(allwordsslice, word2slice[i])
		s := len(word1)
		s2 := len(word2)

		if len(allwordsslice)/2 == len(word1) {
			allwordsslice = append(allwordsslice, word2slice[s:]...)
			return string(allwordsslice)
		} else if len(allwordsslice)/2 == len(word2) {
			allwordsslice = append(allwordsslice, word1slice[s2:]...)
			return string(allwordsslice)
		}

	}

	return string(allwordsslice)

}

func main() {

	word1, word2 := "abcd", "pq"

	fmt.Println(mergeAlternately(word1, word2))

}
