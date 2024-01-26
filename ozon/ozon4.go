package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func test(number int, str string) []int {
	mmax := 30
	mmin := 15
	resultSlices := []int{}
	// number = 25
	if mmax == number && str == ">=" {
		resultSlices = append(resultSlices, mmax)
		return resultSlices
	} else if mmin == number && str == "<=" {
		resultSlices = append(resultSlices, mmin)
		return resultSlices
	}

	if number >= mmin && number <= mmax && str == "<=" {
		mmax = number
		if mmax == mmin {
			resultSlices = append(resultSlices, mmax)
		} else if (mmin + 1) >= mmax {
			resultSlices = append(resultSlices, mmax+1)
		}
	} else {
		resultSlices = append(resultSlices, -1)
	}
	if number >= mmin && number <= mmax && str == ">=" {
		mmin = number
		if mmax == mmin {
			resultSlices = append(resultSlices, mmax)
		} else if (mmax - 1) >= mmin {
			resultSlices = append(resultSlices, mmax-1)
		}
	} else {
		resultSlices = append(resultSlices, -1)
	}

	return resultSlices
}

func main() {
	var numOfSets int

	tMax := 30
	tMin := 15

	fmt.Println("Введите количество наборов входных данных")
	fmt.Scan(&numOfSets)
	var tmp int

	reader := bufio.NewReader(os.Stdin)
	result := ""
	for i := 1; i <= numOfSets; i++ {
		fmt.Println("Введите количество людей")
		fmt.Scan(&tmp)
		maxValue := tMax
		minValue := tMin
		for j := 0; j < tmp; j++ {
			// Ввод букв для текущего набора с клавиатуры
			//	fmt.Printf("Введите буквы для набора %d: ", i)
			setLetters, _ := reader.ReadString('\n')
			setLetters = strings.TrimSpace(setLetters)
			reStr := regexp.MustCompile("[^<>=]+")
			resultStr := reStr.ReplaceAllString(setLetters, "")
			reInt := regexp.MustCompile("[^0-9]+")
			resultInt := reInt.ReplaceAllString(setLetters, "")
			number, err := strconv.Atoi(resultInt)
			if err != nil {
				panic(err)
			}
			if resultStr == ">=" {
				if number > minValue {
					minValue = number
				}
			} else if resultStr == "<=" {
				if number < maxValue {
					maxValue = number
				}

			} else {
				if number+1 > minValue {
					minValue = number + 1
				}
				if number-1 < maxValue {
					maxValue = number - 1
				}
			}

			if minValue > maxValue {
				result += "-1\n"
			} else {
				result += fmt.Sprintf("%d\n", maxValue)
			}
		}
		result += "\n"
	}
	fmt.Print(result)
}
