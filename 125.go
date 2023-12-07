package main

import (
	"fmt"
	"strings"
	"unicode"
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
	s = strings.ToLower(s)
	runes := []rune{}
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			runes = append(runes, ch)
		}
	}
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(isPalindrome("0P"))

}
