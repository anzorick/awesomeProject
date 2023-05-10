package main

import (
	"fmt"
)

// Задание 4. Область видимости переменных
// Что нужно сделать
// Напишите программу, в которой будет три функции, попарно использующие
// разные глобальные переменные. Функции должны прибавлять к поданному на вход числу
// глобальную переменную и возвращать результат. Затем вызовите по очереди три функции,
// передавая результат из одной в другую.

var (
	a = 10
	b = 20
	c = 30
)

func addA(x int) int {
	return x + a
}

func addB(x int) int {
	return x + b
}

func addC(x int) int {
	return x + c
}

func main() {

	var num int
	fmt.Print("введите число: ")
	fmt.Scan(&num)

	num = addA(num)
	fmt.Println(num)
	num = addB(num)
	fmt.Println(num)
	num = addC(num)
	fmt.Println(num)
}
