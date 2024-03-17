package main

/* 1) Ввод строки
2) Разбить на операнды и оператор
3) Проверить наличие арабских или римских цифр
4) Если цифры римкие, перевести в арабские
5) Проверить, что числа не прывышают 10 и не меньше 0
6) Посчитать результат
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// мапа с римскими числами
var mapRIMtoARAB = map[string]int{
	"I":        1,
	"II":       2,
	"III":      3,
	"IV":       4,
	"V":        5,
	"VI":       6,
	"VII":      7,
	"VIII":     8,
	"IX":       9,
	"X":        10,
	"XI":       11,
	"XII":      12,
	"XIII":     13,
	"XIV":      14,
	"XV":       15,
	"XVI":      16,
	"XVII":     17,
	"XVIII":    18,
	"XIX":      19,
	"XX":       20,
	"XXI":      21,
	"XXII":     22,
	"XXIII":    23,
	"XXIV":     24,
	"XXV":      25,
	"XXVI":     26,
	"XXVII":    27,
	"XXVIII":   28,
	"XXIX":     29,
	"XXX":      30,
	"XXXI":     31,
	"XXXII":    32,
	"XXXIII":   33,
	"XXXIV":    34,
	"XXXV":     35,
	"XXXVI":    36,
	"XXXVII":   37,
	"XXXVIII":  38,
	"XXXIX":    39,
	"XL":       40,
	"XLI":      41,
	"XLII":     42,
	"XLIII":    43,
	"XLIV":     44,
	"XLV":      45,
	"XLVI":     46,
	"XLVII":    47,
	"XLVIII":   48,
	"XLIX":     49,
	"L":        50,
	"LI":       51,
	"LII":      52,
	"LIII":     53,
	"LIV":      54,
	"LV":       55,
	"LVI":      56,
	"LVII":     57,
	"LVIII":    58,
	"LIX":      59,
	"LX":       60,
	"LXI":      61,
	"LXII":     62,
	"LXIII":    63,
	"LXIV":     64,
	"LXV":      65,
	"LXVI":     66,
	"LXVII":    67,
	"LXVIII":   68,
	"LXIX":     69,
	"LXX":      70,
	"LXXI":     71,
	"LXXII":    72,
	"LXXIII":   73,
	"LXXIV":    74,
	"LXXV":     75,
	"LXXVI":    76,
	"LXXVII":   77,
	"LXXVIII":  78,
	"LXXIX":    79,
	"LXXX":     80,
	"LXXXI":    81,
	"LXXXII":   82,
	"LXXXIII":  83,
	"LXXXIV":   84,
	"LXXXV":    85,
	"LXXXVI":   86,
	"LXXXVII":  87,
	"LXXXVIII": 88,
	"LXXXIX":   89,
	"XC":       90,
	"XCI":      91,
	"XCII":     92,
	"XCIII":    93,
	"XCIV":     94,
	"XCV":      95,
	"XCVI":     96,
	"XCVII":    97,
	"XCVIII":   98,
	"XCIX":     99,
	"C":        100,
}

// проверяем сущестование римских чисел
func RimToArabExist(x string) bool {
	_, exists := mapRIMtoARAB[x]
	return exists
}

// функция преобразования римского в арабское
func RimToArab(x string) int {
	return mapRIMtoARAB[x]
}

// фунция преобразования арабского в римское
func ArabToRim(value int) string {
	for key, val := range mapRIMtoARAB {
		if val == value {
			return key
		}
	}
	return ""
}

func main() {
	var arg1, arg3 int
	var arg2 string
	var result int
	rim := false  // вводятся римские
	arab := false // вводятся арабские

	fmt.Println("Введите выражение:")

	// ввод строки
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// удляем символ из конца введенной строки
	input = strings.TrimSpace(input)

	//разбиваем элементы по пробелу
	elements := strings.Split(input, " ")
	if len(elements) != 3 {
		panic("Ожидалось два операнда и один оператор")
	}

	// проверяем ввод арабских или римских цифр
	for _, elem := range elements {
		if RimToArabExist(elem) {
			rim = true
		} else {
			if _, err := strconv.Atoi(elem); err == nil {
				arab = true
			}
		}
	}
	// паника, если арабские и римские
	if rim && arab {
		panic("В выражении не могут присутсвовать римские и арабские цирфы одновременно")
	}

	// переводим операнды в числа и проверяем, являются ли они римскими
	if rim {
		arg1 = RimToArab(elements[0])
		arg3 = RimToArab(elements[2])
	} else {
		arg1, err = strconv.Atoi(elements[0])
		if err != nil {
			panic("Ошибка ввода данных: Неверный формат числа")
		}
		arg3, err = strconv.Atoi(elements[2])
		if err != nil {
			panic("Ошибка ввода данных: Неверный формат числа")
		}
	}
	// определяем оператор
	arg2 = elements[1]

	// проверяем деление на 0
	if arg2 == "/" && arg3 == 0 {
		panic("На ноль делить нельзя!")
	}

	// выполняем операцию
	switch arg2 {
	case "+":
		result = arg1 + arg3
	case "-":
		result = arg1 - arg3
	case "*":
		result = arg1 * arg3
	case "/":
		result = arg1 / arg3
	default:
		panic("Неподдерживаемый оператор")
	}

	// проверяем, что оба числа меньше или равны 10 и больше или равны 1
	if (arg1 < 1 || arg1 > 10) || (arg3 < 1 || arg3 > 10) {
		panic("Введенные операнды не входят в допустимый диапазон")
	}
	//выводим результат (с переводом в римскую систему, если операнды изначально римские)
	if rim {
		if result < 1 {
			panic("Результат меньше единицы")
		}
		fmt.Println(ArabToRim(result))
	} else {
		fmt.Println(result)
	}
}
