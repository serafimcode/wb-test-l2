package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//str := "qwe\\\\5"
	str := "qwe\\\\5"
	if unpackedStr, err := unpackString(str); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Initial str: %s, Unpacked str: %s\n", str, unpackedStr)
	}
}

func unpackString(str string) (string, error) {
	/*Для эффективной конкатенации используем strings.Builder*/
	var res strings.Builder

	/*Для удобства преобразуем в руны*/
	runes := []rune(str)

	/*Сразу проверяем пустую строку*/
	if len(runes) == 0 {
		return "", nil
	}
	/*Сразу возвращаем ошибку, если строка начинается с числа*/
	if unicode.IsDigit(runes[0]) {
		return "", errors.New("invalid string")
	}

	for i := 0; i < len(runes); i++ {
		if runes[i] == '\\' {
			/*Если escape последовательность, вставляем следующий символ
			и инкрементируем i, чтобы пропустить следующую итерацию.
			*/
			if i+1 < len(runes) {
				res.WriteString(string(runes[i+1]))
				i++
			} else {
				/*Делаем проверку для последнего в строе символа*/
				res.WriteString(string(runes[i]))
			}
		} else {
			if unicode.IsDigit(runes[i]) {
				/*Если число, вставляем предыдущий символ mult-1,
				т.к. один раз его вставили на прошлом шаге*/
				mult, _ := strconv.Atoi(string(runes[i]))
				res.WriteString(strings.Repeat(string(runes[i-1]), mult-1))
			} else {
				/*Буквы вставляем как есть*/
				res.WriteString(string(runes[i]))
			}
		}
	}

	return res.String(), nil
}
