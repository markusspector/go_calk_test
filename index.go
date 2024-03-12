package main

import (
	"fmt"
	"strconv"
)

func main() {
	var operands []string
	var operator string
	fmt.Print("Введите выражение: ")
	fmt.Scanf("%s", &operator)

	// Читаем операнды и оператор до тех пор, пока не получим пустую строку
	for {
		var operand string
		fmt.Scanf("%s", &operand)
		if operand == "" {
			break
		}
		operands = append(operands, operand)
	}

	// Проверяем операнды на соответствие условиям
	for _, operand := range operands {
		num, err := strconv.Atoi(operand)
		if err != nil && !isRoman(operand) {
			panic("Числа должны быть арабскими или римскими")
		} else if num >= 10 {
			panic("Числа должны быть меньше 10")
		}
	}

	// Преобразуем римские числа в арабские и выполняем вычисление
	nums := make([]int, len(operands))
	for i, operand := range operands {
		num, err := strconv.Atoi(operand)
		if err != nil {
			num = romanToArabic(operand)
		}
		nums[i] = num
	}

	result := calculate(nums, operator)

	// Если один из операндов был римским, выводим результат в римской форме
	if isRoman(operands[0]) {
		fmt.Println("Результат:", arabicToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}
}

func calculate(nums []int, operation string) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		switch operation {
		case "+":
			result += nums[i]
		case "-":
			result -= nums[i]
		case "*":
			result *= nums[i]
		case "/":
			if nums[i] == 0 {
				panic("Деление на 0")
			}
			result /= nums[i]
		default:
			panic("err")
		}
	}
	return result
}

func isRoman(num string) bool {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, romanNumeral := range romanNumerals {
		if num == romanNumeral {
			return true
		}
	}
	return false
}

func romanToArabic(roman string) int {
	switch roman {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		panic("err")
	}
}

func arabicToRoman(arabic int) string {
	switch arabic {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	case 10:
		return "X"
	default:
		panic("err")
	}
}
