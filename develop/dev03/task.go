package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Утилита sort
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/

// Определяем флаги
var (
	columnFlag  = flag.Int("k", -1, "Sort by column")
	numericFlag = flag.Bool("n", false, "Sort in order of string numericFlag value")
	dedupeFlag  = flag.Bool("u", false, "Remove duplicates")
	reverseFlag = flag.Bool("r", false, "Sort in reversed order")
)

func main() {
	// Парсим флаги
	flag.Parse()
	//Проверяем на корректность вызова программы
	if flag.NArg() != 1 {
		fmt.Println("Usage: sort.go [-r] <filename>")
		os.Exit(1)
	}

	// Берем имя файла и открываем файл
	filename := flag.Arg(0)
	lines, err := read(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Сортируем
	if *numericFlag {
		lines, err = sortNumeric(lines)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	} else if *columnFlag >= 0 {
		lines, err = sortColumns(lines, *columnFlag)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	} else {
		sort.Strings(lines)
	}

	// Форматируем вывод
	if *reverseFlag {
		lines = reverse(lines)
	}
	if *dedupeFlag {
		lines = dedupe(lines)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func read(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("error opening file")
	}
	defer file.Close()

	// Читаем данные
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Проверяем сканер на наличие ошибок во время чтения
	if err := scanner.Err(); err != nil {
		return nil, errors.New("error reading file")
	}

	return lines, nil
}

func sortNumeric(lines []string) ([]string, error) {
	var res []string
	var numericValues []int
	var tempStr string

	for _, l := range lines {
		tempStr = strings.TrimSuffix(l, "\n")
		tempStr = strings.Replace(tempStr, "\r", "", 1)

		v, _ := strconv.Atoi(l)
		numericValues = append(numericValues, v)
	}
	sort.Ints(numericValues)
	for _, v := range numericValues {
		res = append(res, strconv.Itoa(v))
	}

	return res, nil
}

/*Сортируем по колонкам*/
func sortColumns(lines []string, column int) ([]string, error) {
	var strByColumns [][]string
	var res []string
	/*Приводим данные, введенные пользователем к индексам*/
	column -= 1

	/*Разбиваем строки по пробелам и возвращаем ошибку,
	количество слов в строке меньше номера столбца,
	указанного пользователем*/
	for _, l := range lines {
		values := strings.Split(l, " ")
		if len(values)-1 < column {
			return nil, fmt.Errorf("there is no such column = %d", column)
		}
		strByColumns = append(strByColumns, values)
	}

	/*Сортируем строки по номеру колонки*/
	sort.Slice(strByColumns, func(i, j int) bool {
		if len(strByColumns[i]) == 0 || len(strByColumns[j]) == 0 {
			return len(strByColumns[i]) == 0
		}

		return strByColumns[i][column] < strByColumns[j][column]
	})
	for _, s := range strByColumns {
		res = append(res, strings.Join(s, " "))
	}

	return res, nil
}

/*Сортируем в обратном порядке*/
func reverse(lines []string) []string {
	res := make([]string, len(lines))
	lstIdx := len(lines) - 1

	for i := range lines {
		res[i] = lines[lstIdx-i]
	}

	return res
}

/*Убираем повторы*/
func dedupe(lines []string) []string {
	set := make(map[string]struct{})
	var res []string

	for _, l := range lines {
		/*Если в мапе еще нет записи, добавляем строку
		и делаем запись в мапу, чтобы в следующий раз
		пропустить строку, если будет повтор*/
		if _, found := set[l]; !found {
			res = append(res, l)
			set[l] = struct{}{}
		}
	}
	return res
}
