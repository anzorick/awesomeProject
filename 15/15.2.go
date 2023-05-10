package main

import "fmt"

func main() {
	var a [5]int = [5]int{10, 20, 30, 40, 50}
	// a := [...]int {10, 20, 30, 40, 50} аналог
	a[2] = 70
	fmt.Println(a)
}
