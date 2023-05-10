//Разработайте программу, позволяющую ввести 10 целых чисел, а затем вывести из них количество
//чётных и нечётных чисел. Для ввода и подсчёта используйте разные циклы.
//
//Что оценивается
//Для введённых чисел 1, 1, 1, 2, 2, 2, 3, 3, 3, 4 программа должна вывести: чётных — 4, нечётных — 6.

package main

import "fmt"

func main() {
	var nums [10]int

	for i := 0; i < 10; i++ {
		fmt.Print("введите числа: ")
		fmt.Scan(&nums[i])
	}

	slice1 := []int{}
	slice2 := []int{}
	for _, n := range nums {
		if n%2 == 0 {
			slice1 = append(slice1, n)
		} else {
			slice2 = append(slice2, n)
		}
	}
	fmt.Print("эти числа четные: ")
	for _, s1 := range slice1 {
		fmt.Print(s1, " ")
	}
	fmt.Println()

	fmt.Print("эти числа нечетные: ")
	for _, s2 := range slice2 {
		fmt.Print(s2, " ")
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("эти числа четные: ", slice1)
	fmt.Println("эти числа нечетные: ", slice2)
}
