package main

import "strings"

func reverseVowels(s string) string {
	i := 0
	j := len(s) - 1
	vowels := "aeiouAEIOU"
	result := make([]byte, len(s))
	for i <= j {
		if !strings.Contains(vowels, string(s[i])) {
			result[i] = s[i]
			i++
			continue
		}
		if !strings.Contains(vowels, string(s[j])) {
			result[j] = s[j]
			j--
			continue
		}

		result[i], result[j] = s[j], s[i]
		i++
		j--
	}

	return string(result[:])
}
