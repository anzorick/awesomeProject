package main

import "fmt"

func main() {
	var arr [100]int
	var n int
	fmt.Println("введите количество элементов")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("arr[%v]=", i)
		fmt.Scan(&arr[i])
	}
	fmt.Println("вывод")
	for i := 0; i < (n+1)/2; i++ { // (n+1) для того чтобы выводилось коректное количество элементов при запросе
		index := i*2 + 1
		fmt.Printf("arr[%v]=%v\n", index, arr[index])
	}
}
