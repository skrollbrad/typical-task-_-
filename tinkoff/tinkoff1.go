package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Строка TINKOFF
	target := "1111222334"
	var sl []string
	// Ввод строки с клавиатуры
	reader := bufio.NewReader(os.Stdin)

	// Ввод количества наборов с клавиатуры
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Введите количество наборов букв: ")
	var numOfSets int
	fmt.Scan(&numOfSets)

	// Проверка для каждого набора букв
	for i := 1; i <= numOfSets && numOfSets > 0 && numOfSets < 101; i++ {
		// Ввод букв для текущего набора с клавиатуры
		fmt.Printf("Введите буквы для набора %d: ", i)
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

		// Проверка, что в наборе есть все буквы из строки "TINKOFF" с нужными количествами
		canFormTarget := true
		for _, ch := range target {
			if letterCounts[ch] < strings.Count(target, string(ch)) {
				canFormTarget = false
				break
			}
		}

		// Вывод результата для текущего набора
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
