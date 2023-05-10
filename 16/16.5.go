package main

import "fmt"

func main() {
	fmt.Println("Введите ширину и высоту поля")
	var a, b int
	fmt.Scan(&a, &b)

	for k := 0; k < b; k++ {
		cycle2(a, k)
		fmt.Println()
	}
	fmt.Println("end")
}

func cycle2(a int, k int) {
	for j := 0; j < a; j++ {
		if (k+j)%2 == 0 {
			fmt.Print("*")
		} else {
			fmt.Print(" ")
		}
	}
}
