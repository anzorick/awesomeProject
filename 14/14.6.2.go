package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Задание 2. Функция, возвращающая несколько значений
// Что нужно сделать
// Напишите программу, которая с помощью функции генерирует три случайные точки
// в двумерном пространстве (две координаты), а затем с помощью другой функции преобразует
// эти координаты по формулам: x = 2 × x + 10, y = −3 × y − 5.

func randomPoint() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100), rand.Intn(100)
}

func transformPoint(x, y int) (int, int) {
	x = 2*x + 10
	y = -3*y - 5
	return x, y
}

func main() {
	x1, y1 := randomPoint()
	x2, y2 := randomPoint()
	x3, y3 := randomPoint()

	fmt.Println("Исходные точки:")
	fmt.Println("Точка 1:", x1, y1)
	fmt.Println("Точка 2:", x2, y2)
	fmt.Println("Точка 3:", x3, y3)

	x1, y1 = transformPoint(x1, y1)
	x2, y2 = transformPoint(x2, y2)
	x3, y3 = transformPoint(x3, y3)

	fmt.Println("Точки после преобразования:")
	fmt.Println("Точка 1:", x1, y1)
	fmt.Println("Точка 2:", x2, y2)
	fmt.Println("Точка 3:", x3, y3)
}
