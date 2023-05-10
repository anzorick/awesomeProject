package main

import (
	"fmt"
	"strconv"
)

//Задание 1. Функция, возвращающая результат
//Что нужно сделать
//Напишите функцию, которая на вход получает число и возвращает true, если число четное, и false, если нечётное.
//Рекомендация
//Программа запрашивает у пользователя или генерирует случайное число,
//передает в функцию в качестве аргумента и выводит в консоль результат её работы.

func number(a int) bool {
	return a%2 == 0
}

func main() {
	var input string
	for {
		fmt.Println("Введите число")
		fmt.Scanln(&input)

		if input == "stop" {
			fmt.Println("Завершение работы программы")
			return
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		if number(num) {
			fmt.Println(num, "четное")
		} else {
			fmt.Println(num, "нечетное")
		}

	}

}
