package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	table = [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C"},
	}
)

type Operand struct {
	value int
	roman bool
}

func arab2roman(arab int) string {
	if arab <= 0 {
		panic("В римской системе нет отрицательных чисел.")
	}

	var result string
	digit := 100
	for i := 2; i >= 0; i-- {
		d := arab / digit
		result += table[i][d]
		arab %= digit
		digit /= 10
	}

	return result
}

func parseOperand(s string) Operand {
	re := regexp.MustCompile(`^(L?X{0,3}|X[LC])(V?I{0,3}|I[VX])$`)

	result := Operand{}
	result.roman = false

	val, err := strconv.Atoi(s)
	if val <= 0 && val > 10 {
		panic("Операнд должен быть в диапазоне от 1 до 10 или от I до X.")
	}
	if err != nil {
		if re.MatchString(s) {
			digits := re.FindAllStringSubmatch(s, 1)

			arab := 0
			digit := 10
			for i := 1; i < len(digits[0]); i++ {
				for k, v := range table[2-i] {
					if digits[0][i] == v {
						arab += k * digit
						break
					}
				}
				digit /= 10
			}

			val = arab
			result.roman = true
		} else {
			panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}
	}

	result.value = val
	return result
}

func calculate(o1 Operand, o2 Operand, operator string) {
	if o1.roman != o2.roman {
		panic("Используются одновременно разные системы счисления.")
	}
	val := 0

	switch operator {
	case "+":
		val = o1.value + o2.value
	case "-":
		if o1.value < o2.value {
			panic("В римской системе нет отрицательных чисел.")
		}
		val = o1.value - o2.value
	case "/":
		val = o1.value / o2.value
	case "*":
		val = o1.value * o2.value
	default:
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	result := fmt.Sprint(val)

	if o1.roman {
		result = arab2roman(val)
	}

	fmt.Println(result)
}

func main() {
	fmt.Println("Калькулятор")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите выражение (Например: 2 + 7):")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ToUpper(text)
		fields := strings.Fields(text)

		if len(fields) == 1 {
			panic("Строка не является математической операцией.")
		}
		if len(fields) != 3 {
			panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}

		calculate(parseOperand(fields[0]), parseOperand(fields[2]), fields[1])
	}
}
