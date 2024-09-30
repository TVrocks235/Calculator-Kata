package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Конвертация римских чисел в арабские
var romanToArabicMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Конвертация арабских чисел в римские
var arabicToRomanMap = []struct {
	Arabic int
	Roman  string
}{
	{10, "X"}, {9, "IX"}, {8, "VIII"}, {7, "VII"},
	{6, "VI"}, {5, "V"}, {4, "IV"}, {3, "III"}, {2, "II"}, {1, "I"},
}

// Функция конвертации римского числа в арабское
func romanToArabic(roman string) (int, error) {
	if value, ok := romanToArabicMap[roman]; ok {
		return value, nil
	}
	return 0, errors.New("неверный формат римского числа")
}

// Функция конвертации арабского числа в римское
func arabicToRoman(num int) (string, error) {
	if num < 1 {
		return "", errors.New("результат римского числа должен быть больше или равен 1")
	}
	var result strings.Builder
	for _, pair := range arabicToRomanMap {
		for num >= pair.Arabic {
			result.WriteString(pair.Roman)
			num -= pair.Arabic
		}
	}
	return result.String(), nil
}

// Функция выполнения арифметической операции
func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль невозможно")
		}
		return a / b, nil
	default:
		return 0, errors.New("неверный оператор")
	}
}

// Функция для обработки ввода
func processInput(input string) (string, error) {
	// Определение, римские ли числа или арабские
	isRoman := false
	for _, r := range input {
		if strings.Contains("IVXLCDM", string(r)) {
			isRoman = true
			break
		}
	}

	// Поиск оператора
	var operator string
	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	} else if strings.Contains(input, "*") {
		operator = "*"
	} else if strings.Contains(input, "/") {
		operator = "/"
	} else {
		return "", errors.New("оператор не найден")
	}

	// Разделение строки на операнды
	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		return "", errors.New("неверный формат выражения")
	}
	aStr := strings.TrimSpace(parts[0])
	bStr := strings.TrimSpace(parts[1])

	var a, b int
	var err error

	// Обработка римских чисел
	if isRoman {
		a, err = romanToArabic(aStr)
		if err != nil {
			return "", err
		}
		b, err = romanToArabic(bStr)
		if err != nil {
			return "", err
		}
	} else {
		// Обработка арабских чисел
		a, err = strconv.Atoi(aStr)
		if err != nil {
			return "", errors.New("неверный формат арабского числа")
		}
		b, err = strconv.Atoi(bStr)
		if err != nil {
			return "", errors.New("неверный формат арабского числа")
		}
	}

	// Проверка диапазона
	if a < 1 || a > 10 || b < 1 || b > 10 {
		return "", errors.New("числа должны быть в диапазоне от 1 до 10")
	}

	// Выполнение вычислений
	result, err := calculate(a, b, operator)
	if err != nil {
		return "", err
	}

	// Возврат результата
	if isRoman {
		if result < 1 {
			return "", errors.New("результат римского числа должен быть больше или равен 1")
		}
		return arabicToRoman(result)
	}

	return strconv.Itoa(result), nil
}

func main() {
	// Ввод данных пользователем
	var input string
	fmt.Println("Введите выражение (например: 3 + 5 или IV * II):")
	fmt.Scanln(&input)

	// Обработка и вывод результата
	result, err := processInput(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", result)

	fmt.Println("Программа завершена. Нажмите Enter, чтобы выйти...")
	fmt.Scanln() // Ожидает ввода от пользователя
}
