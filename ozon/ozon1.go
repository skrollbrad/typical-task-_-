package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Строка исходня
	target := "1111222334"
	var sl []string
	// Ввод строки с клавиатуры

	// Ввод количества наборов с клавиатуры

	//fmt.Print("Введите количество наборов букв: ")
	var numOfSets int
	fmt.Scan(&numOfSets)
	reader := bufio.NewReader(os.Stdin)

	// Проверка для каждого набора букв
	for i := 1; i <= numOfSets && numOfSets > 0 && numOfSets < 1001; i++ {
		// Ввод букв для текущего набора с клавиатуры
		//	fmt.Printf("Введите буквы для набора %d: ", i)
		setLetters, _ := reader.ReadString('\n')
		setLetters = strings.TrimSpace(setLetters)
		if len(setLetters) > 20 {
			continue
		}

		// Создаем карту, содержащую количество каждой буквы в наборе
		letterCounts := make(map[rune]int)
		for _, ch := range setLetters {
			letterCounts[ch]++
		}

		// Проверка, что в наборе есть все цифры из строки "target" с нужными количествами
		canFormTarget := true
		for _, ch := range target {
			if letterCounts[ch] < strings.Count(target, string(ch)) {
				canFormTarget = false
				break
			}
		}

		if canFormTarget {
			sl = append(sl, "YES")
		} else {
			sl = append(sl, "NO")
		}

	}
	for _, v := range sl {
		fmt.Println(v)
	}

}
