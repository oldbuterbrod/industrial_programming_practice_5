package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Линейное программирование
// Перевод чисел из одной системы счисления в другую
func convertBase(number string, fromBase, toBase int) (string, error) {
	decimal, err := strconv.ParseInt(number, fromBase, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(decimal, toBase), nil
}

// Решение квадратного уравнения
func solveQuadratic(a, b, c float64) (complex128, complex128) {
	discriminant := b*b - 4*a*c
	if discriminant >= 0 {
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return complex(root1, 0), complex(root2, 0)
	}
	realPart := -b / (2 * a)
	imagPart := math.Sqrt(-discriminant) / (2 * a)
	return complex(realPart, imagPart), complex(realPart, -imagPart)
}

// Сортировка чисел по модулю
func sortByAbs(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return math.Abs(float64(arr[i])) < math.Abs(float64(arr[j]))
	})
	return arr
}

// Слияние двух отсортированных массивов
func mergeSortedArrays(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)
	return result
}

// Нахождение подстроки в строке без использования встроенных функций
func findSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		j := 0
		for j < len(substr) && s[i+j] == substr[j] {
			j++
		}
		if j == len(substr) {
			return i
		}
	}
	return -1
}

// Условия
// Калькулятор с расширенными операциями
func calculate(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	case "^":
		return math.Pow(a, b), nil
	case "%":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return float64(int(a) % int(b)), nil
	default:
		return 0, errors.New("недопустимая операция")
	}
}

// Проверка палиндрома
func isPalindrome(s string) bool {
	var cleaned []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}
	for i := 0; i < len(cleaned)/2; i++ {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
			return false
		}
	}
	return true
}

// Нахождение пересечения трех отрезков
func intersectionExists(a1, a2, b1, b2, c1, c2 int) bool {
	left := max(a1, max(b1, c1))
	right := min(a2, min(b2, c2))
	return left <= right
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Выбор самого длинного слова в предложении
func longestWord(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	var longest string
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

// Проверка високосного года
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// Циклы
// Числа Фибоначчи до определенного числа
func fibonacciUpTo(n int) []int {
	var fib []int
	a, b := 0, 1
	for a <= n {
		fib = append(fib, a)
		a, b = b, a+b
	}
	return fib
}

// Определение простых чисел в диапазоне
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primesInRange(start, end int) []int {
	var primes []int
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// Числа Армстронга в заданном диапазоне
func isArmstrong(n int) bool {
	sum := 0
	temp := n
	digits := int(math.Log10(float64(n)) + 1)
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

func armstrongInRange(start, end int) []int {
	var armstrongs []int
	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			armstrongs = append(armstrongs, i)
		}
	}
	return armstrongs
}

// Реверс строки
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Нахождение наибольшего общего делителя (НОД)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Меню выбора задач
func main() {
	for {
		fmt.Println("\nВыберите задачу:")
		fmt.Println("1. Перевод чисел из одной системы счисления в другую")
		fmt.Println("2. Решение квадратного уравнения")
		fmt.Println("3. Сортировка чисел по модулю")
		fmt.Println("4. Слияние двух отсортированных массивов")
		fmt.Println("5. Нахождение подстроки в строке")
		fmt.Println("6. Калькулятор")
		fmt.Println("7. Проверка палиндрома")
		fmt.Println("8. Пересечение трех отрезков")
		fmt.Println("9. Самое длинное слово")
		fmt.Println("10. Проверка високосного года")
		fmt.Println("11. Числа Фибоначчи")
		fmt.Println("12. Простые числа в диапазоне")
		fmt.Println("13. Числа Армстронга")
		fmt.Println("14. Реверс строки")
		fmt.Println("15. НОД")
		fmt.Println("0. Выход")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var number string
			var fromBase, toBase int
			fmt.Println("Введите число, исходную и конечную систему счисления:")
			fmt.Scan(&number, &fromBase, &toBase)
			result, err := convertBase(number, fromBase, toBase)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Результат:", result)
			}
		case 2:
			var a, b, c float64
			fmt.Println("Введите коэффициенты a, b, c:")
			fmt.Scan(&a, &b, &c)
			root1, root2 := solveQuadratic(a, b, c)
			fmt.Printf("Корни уравнения: %v и %v\n", root1, root2)
		case 3:
			arr := []int{-7, -2, 3, 9, -4}
			fmt.Println("Исходный массив:", arr)
			sortedArr := sortByAbs(arr)
			fmt.Println("Отсортированный по модулю:", sortedArr)
		case 4:
			arr1 := []int{1, 3, 5}
			arr2 := []int{2, 4, 6}
			mergedArr := mergeSortedArrays(arr1, arr2)
			fmt.Println("Объединенный массив:", mergedArr)
		case 5:
			var s, substr string
			fmt.Println("Введите строку и подстроку для поиска:")
			fmt.Scan(&s, &substr)
			pos := findSubstring(s, substr)
			fmt.Println("Индекс первого вхождения:", pos)
		case 6:
			var a, b float64
			var operator string
			fmt.Println("Введите два числа и оператор (+, -, *, /, ^, %):")
			fmt.Scan(&a, &b, &operator)
			result, err := calculate(a, b, operator)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Результат:", result)
			}
		case 7:
			var s string
			fmt.Println("Введите строку для проверки палиндрома:")
			fmt.Scan(&s)
			fmt.Println("Палиндром?", isPalindrome(s))
		case 8:
			var a1, a2, b1, b2, c1, c2 int
			fmt.Println("Введите три отрезка (начало и конец каждого):")
			fmt.Scan(&a1, &a2, &b1, &b2, &c1, &c2)
			fmt.Println("Пересечение существует?", intersectionExists(a1, a2, b1, b2, c1, c2))
		case 9:
			var sentence string
			fmt.Println("Введите предложение:")
			fmt.Scan(&sentence)
			fmt.Println("Самое длинное слово:", longestWord(sentence))
		case 10:
			var year int
			fmt.Println("Введите год:")
			fmt.Scan(&year)
			fmt.Println("Високосный год?", isLeapYear(year))
		case 11:
			var n int
			fmt.Println("Введите максимальное значение Фибоначчи:")
			fmt.Scan(&n)
			fmt.Println("Числа Фибоначчи:", fibonacciUpTo(n))
		case 12:
			var start, end int
			fmt.Println("Введите диапазон для поиска простых чисел:")
			fmt.Scan(&start, &end)
			fmt.Println("Простые числа в диапазоне:", primesInRange(start, end))
		case 13:
			var start, end int
			fmt.Println("Введите диапазон для поиска чисел Армстронга:")
			fmt.Scan(&start, &end)
			fmt.Println("Числа Армстронга:", armstrongInRange(start, end))
		case 14:
			var s string
			fmt.Println("Введите строку для реверса:")
			fmt.Scan(&s)
			fmt.Println("Перевернутая строка:", reverseString(s))
		case 15:
			var a, b int
			fmt.Println("Введите два числа для нахождения НОД:")
			fmt.Scan(&a, &b)
			fmt.Println("НОД:", gcd(a, b))
		case 0:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор. Пожалуйста, попробуйте снова.")
		}
	}
}
