package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Фраза является палиндромом , если после преобразования всех заглавных букв в
// строчные и удаления всех небуквенно-цифровых символов она читается одинаково и вперед,
// и назад. Буквенно-цифровые символы включают буквы и цифры.

// Учитывая строку s, возвращаться true если это палиндром , или falseв противном случае .

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

func isPalindrome(s string) bool {

	newstr := strings.ToLower(s)
	newstr = strings.ReplaceAll(newstr, " ", "")
	reg := regexp.MustCompile("[^a-zA-Z]+")
	result := reg.ReplaceAllString(newstr, "")
	runes := []rune(result)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	runes2 := []rune(result)

	fmt.Println(runes)
	fmt.Println(runes2)
	if reflect.DeepEqual(runes, runes2) {
		return true

	}

	fmt.Println(runes)
	fmt.Println(runes2)

	return false

}

func main() {

	fmt.Println(isPalindrome("0P"))

}
