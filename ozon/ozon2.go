package main

import "fmt"

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func isValidDate(day, month, year int) bool {
	if year < 1950 || year > 2300 {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	if day < 1 {
		return false
	}
	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeapYear(year) {
		daysInMonth[1] = 29
	}
	return day <= daysInMonth[month-1]
}

func main() {
	var t, day, month, year int
	slice := []string{}
	fmt.Scanln(&t) // Вводим количество наборов входных данных
	for i := 0; i < t; i++ {
		fmt.Scanln(&day, &month, &year) // Вводим день, месяц и год
		if isValidDate(day, month, year) {
			slice = append(slice, "YES")
		} else {
			slice = append(slice, "NO")
		}
	}
	for _, i := range slice {
		fmt.Println(i)
	}
}
