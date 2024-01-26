// либо автомобильный номер имеет вид
// буква−цифра−цифра−буква−буква (примеры корректных номеров первого вида:R48FA R48FA, O00OO O00OO, A99OK A99OK);
// либо автомобильный номер имеет вид
// буква−цифра−буква−буква (примеры корректных номеров второго вида:T7RR T7RR, A9PQ A9PQ, O0OO O0OO).
// R48FA O00OOO0OOA99OKA99O
// если i-символ меньше кода буквы/цифры, то прибавляем +1 символ к строке, и если опять, то -, иначе берем дальше
// нужно смотреть первые 4 и первые 5 и если находим удалять крч
// можно в мапу засовывать и удалять ко key
package main

import (
	"fmt"
	"unicode"
)

func isValidNumber(s string) bool {
	if len(s) == 5 {
		for i, char := range s {
			if i == 1 || i == 2 {
				if !unicode.IsDigit(char) {
					return false
				}
			} else {
				if !unicode.IsLetter(char) || !unicode.IsUpper(char) {
					return false
				}
			}
		}
		return true
	} else if len(s) == 4 {
		for i, char := range s {
			if i == 1 {
				if !unicode.IsDigit(char) {
					return false
				}
			} else {
				if !unicode.IsLetter(char) || !unicode.IsUpper(char) {
					return false
				}
			}
		}
		return true
	} else {
		return false
	}
}

func main() {
	var n int
	var s string
	var results []string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		result := ""
		valid := true
		for len(s) > 0 {
			if len(s) >= 5 && isValidNumber(s[:5]) {
				result += s[:5] + " "
				s = s[5:]
			} else if len(s) >= 4 && isValidNumber(s[:4]) {
				result += s[:4] + " "
				s = s[4:]
			} else {
				valid = false
				break
			}
		}
		if valid {
			results = append(results, result)
		} else {
			results = append(results, "-")
		}
	}
	for _, res := range results {
		fmt.Println(res)
	}
}
