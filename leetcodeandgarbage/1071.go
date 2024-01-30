package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}

	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}

	if str1[:len(str2)] != str2 {
		return ""
	}

	return gcdOfStrings(str1[len(str2):], str2)
}
func main() {

	str1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	str2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"

	fmt.Println(gcdOfStrings(str1, str2))

}
