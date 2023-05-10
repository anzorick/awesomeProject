package main

import "fmt"

// Задание 3. Именованные возвращаемые значения
// Что нужно сделать
// Напишите программу, которая на вход получает число, затем с помощью двух функций
// преобразует его. Первая умножает, а вторая прибавляет число, используя именованные возвращаемые значения.
func multiply(number int) int {
	return number * 10
}

func sum(number int) int {
	return number + 10
}

func main() {
	var number int
	fmt.Print("введите число: ")
	fmt.Scan(&number)

	result := multiply(number)
	fmt.Println(result)
	number = sum(number)
	fmt.Println(number)

}
