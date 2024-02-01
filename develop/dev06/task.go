package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*

Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

type config struct {
	fields         int
	delimiter      string
	hideWrongLines bool
}

func parseFlags() config {
	cfg := config{}

	flag.IntVar(&cfg.fields, "f", 0, "показать все столбцы или только этот столбец")
	flag.StringVar(&cfg.delimiter, "d", "\t", "разделитель")
	flag.BoolVar(&cfg.hideWrongLines, "s", true, "скрывать ли строки без разделителя")
	flag.Parse()

	return cfg
}

func cut(sc *bufio.Scanner, cfg config) {

	for sc.Scan() {
		line := sc.Text()
		if !cfg.hideWrongLines && !strings.Contains(line, cfg.delimiter) {
			continue
		}
		columns := strings.Split(line, cfg.delimiter)

		if cfg.fields == 0 {
			fmt.Println(strings.Join(columns, " "))
		} else {
			if len(columns) >= cfg.fields {
				fmt.Println(columns[cfg.fields-1])
			}
		}
	}
}

func main() {
	cfg := parseFlags()
	sc := bufio.NewScanner(os.Stdin)

	cut(sc, cfg)
}
