package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

var errLogger = log.Default()

func main() {
	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		catchError(err)
	}
	fmt.Println(currentTime)
}

func catchError(err error) {
	errLogger.Println(err)
	os.Exit(1)
}
