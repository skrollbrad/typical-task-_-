package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	var str string
	for i := len(words) - 1; i >= 0; i-- {
		str += words[i] + " "

	}
	newStr := str[0 : len(str)-1]

	return string(newStr)
}
func main() {

	str := "hello world"
	fmt.Println(reverseWords(str))

}
