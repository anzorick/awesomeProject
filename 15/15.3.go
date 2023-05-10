package main

import "fmt"

func main() {
	var rain [10]int
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("введите количество осадков в %v году ", 2011+i)
	//	fmt.Scan(&rain[i])
	//}
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("год: %v осадков: %v\n", 2011+i, rain[i])
	//}
	for i, _ := range rain {
		year := 2011 + i
		fmt.Printf("введите количество осадков в %v году ", year)
		fmt.Scan(&rain[i])
	}
	for i, r := range rain {
		fmt.Printf("год: %v осадков: %v\n", 2011+i, r)
	}
}
