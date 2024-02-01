package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.


Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
*/

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := groupAnagrams(words)

	for key, value := range anagrams {
		fmt.Println(key, value)
	}
}

func groupAnagrams(strs []string) map[string][]string {
	keyMp := map[string][]string{}

	for _, s := range strs {
		key := getKey(s)
		keyMp[key] = append(keyMp[key], s)
	}

	/*Меняем ключ, на первый элемент массива*/
	for k, v := range keyMp {
		newKey := v[0]
		val := v
		delete(keyMp, k)

		keyMp[newKey] = val
	}

	return keyMp
}

/*
Из строки делаем ключ, который будет одинаковым
для любой другой строки, с тем же набором символов
*/
func getKey(str string) string {
	intKeys := [33]int{}

	for _, v := range str {
		c := v - 'а'
		intKeys[c] += 1
	}

	sb := strings.Builder{}
	for _, v := range intKeys {
		sb.Write([]byte(strconv.Itoa(v)))
	}
	return sb.String()
}
